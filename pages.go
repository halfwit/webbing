package router

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/olmaxmedical/database"
	"golang.org/x/text/message"
)

var pagecache map[string]*Page

func init() {
	pagecache = make(map[string]*Page)
}

// Access defines the access rights for a specific page
type Access uint8

const (
	// GuestAuth - non registered user
	GuestAuth Access = 1 << iota
	// PatientAuth - normal user, added by registration process
	PatientAuth
	// DoctorAuth - manually added to the database
	DoctorAuth
)

// Page defines what a client receives from a GET request
type Page struct {
	Access Access
	Extra  PluginMask
	CSS    string
	Path   string
	Data   func(p *message.Printer) map[string]interface{}
	tmpl   *template.Template
}

// AddPage - register a *Page to the cache
func AddPage(p *Page) {
	pagecache[p.Path+".html"] = p
}

// ValidatePages - Walk all our templates, test them, and finally return applicable errors as an array
func ValidatePages() []error {
	var errs []error
	hd := path.Join("templates", "header.tpl")
	fd := path.Join("templates", "footer.tpl")
	ld := path.Join("templates", "layout.tpl")
	extra, err := os.Open(path.Join("templates", "plugins"))
	if err != nil {
		errs = append(errs, errors.New("Unable to locate templates/plugins"))
		return errs
	}
	dirs, err := extra.Readdirnames(0)
	for n, dir := range dirs {
		dirs[n] = path.Join("templates", "plugins", dir)
	}
	// TODO(halfwit) Validate our plugin templates here as well
	dirs = append(dirs, hd, fd, ld)
	printer := message.NewPrinter(message.MatchLanguage("en"))
	for _, item := range pagecache {
		var err error
		tp := path.Join("templates", item.Path) + ".tpl"
		t := template.New(path.Base(tp))
		// TODO(halfwit) Contemplate only adding templates for plugins each page uses
		item.tmpl, _ = t.ParseFiles(dirs...)
		item.tmpl, err = t.ParseFiles(tp)
		if err != nil {
			errs = append(errs, fmt.Errorf("parsing in %s - %v", path.Dir(item.Path), err))
			continue
		}
		p := &Request{
			printer: printer,
			path:    item.Path + ".html",
			role:    database.PatientAuth | database.DoctorAuth | database.GuestAuth,
		}
		_, err = getData(p, "")
		if err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

func getpage(p *Request, w http.ResponseWriter) {
	var data []byte
	var err error
	switch database.UserRole(p.user) {
	case database.DoctorAuth:
		data, err = getData(p, "doctor")
	case database.PatientAuth:
		data, err = getData(p, "patient")
	default:
		data, err = getData(p, "guest")
	}
	if err != nil && err.Error() == "Unauthorized" {
		p.Session().Set("redirect", p.path)
		http.Redirect(w, p.Request(), "/login.html", 302)
		return
	}
	if err != nil {
		http.Error(w, "Service Unavailable", 503)
		return
	}
	fmt.Fprintf(w, "%s", data)
}

func getData(p *Request, in string) ([]byte, error) {
	cache, ok := pagecache[p.path]
	if !ok {
		return nil, fmt.Errorf("No such page: %s", p.path)
	}
	if uint8(p.role)&uint8(cache.Access) == 0 {
		return nil, errors.New("Unauthorized")
	}
	r := cache.Data(p.printer)
	r["css"] = cache.CSS
	r["header"] = header(p.printer, p.status)
	r["footer"] = footer(p.printer)
	r["basedir"] = getBaseDir(cache.Path)
	// TODO(halfwit) Test chunking in to go routines if n gets too large
	for _, key := range pluginKey {
		if (cache.Extra&key) != 0 && pluginCache[key].Run != nil {
			r[pluginCache[key].Name] = pluginCache[key].Run(p)
		}
	}
	if p.session != nil {
		r["username"] = p.session.Get("username")
		if _, ok := p.session.Get("redirect").(string); ok {
			r["redirect"] = true
		}
	}
	return cache.render(r)
}

func (page *Page) render(i map[string]interface{}) ([]byte, error) {
	var buf bytes.Buffer
	data := bufio.NewWriter(&buf)
	err := page.tmpl.ExecuteTemplate(data, "layout", i)
	data.Flush()
	return buf.Bytes(), err

}

func getBaseDir(fp string) string {
	return strings.Repeat("../", strings.Count(fp, "/"))
}

package router

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"path"
	"strings"

	"github.com/olmaxmedical/olmax_go/db"
	"golang.org/x/text/message"
)

var pagecache map[string]*Page

//var countrylist []Country

func init() {
	pagecache = make(map[string]*Page)
	//countrylist = listcountries()
}

// Access defines the access rights for a specific page
type Access uint8

const (
	GuestAuth Access = 1 << iota
	PatientAuth
	DoctorAuth
)

// IncludeExtra - helper bitmasks to populate common elements across the site
type IncludeExtra uint32

const (
	FormToken IncludeExtra = 1 << iota
	FormErrors
	SessionToken
)

// Page defines what a client receives from a GET request
type Page struct {
	Access Access
	Extra  IncludeExtra
	CSS    string
	Path   string
	Data   func(p *message.Printer) map[string]interface{}
	tmpl   *template.Template
}

// AddPage - register a *Page to the cache
func AddPage(p *Page) {
	pagecache[p.Path+".html"] = p
}

// ValidatePages - Walk all our templates and finally return applicable errors as an array
func ValidatePages() []error {
	var errs []error
	hd := path.Join("templates", "header.tpl")
	fd := path.Join("templates", "footer.tpl")
	ed := path.Join("templates", "errors.tpl")
	ld := path.Join("templates", "layout.tpl")
	printer := message.NewPrinter(message.MatchLanguage("en"))
	for _, item := range pagecache {
		var err error
		tp := path.Join("templates", item.Path) + ".tpl"

		t := template.New(path.Base(tp))
		item.tmpl, err = t.ParseFiles(tp, hd, ed, fd, ld)
		if err != nil {
			errs = append(errs, fmt.Errorf("parsing in %s - %v", path.Dir(item.Path), err))
			continue
		}
		p := &request{
			printer: printer,
			path:    item.Path + ".html",
			role:    db.PatientAuth | db.DoctorAuth | db.GuestAuth,
		}
		_, err = getdata(p, "")
		if err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

func getdata(p *request, in string) ([]byte, error) {
	cache, ok := pagecache[p.path]
	if !ok {
		return nil, fmt.Errorf("No such page: %s", p.path)
	}
	if uint8(p.role)&uint8(cache.Access) == 0 {
		return nil, fmt.Errorf("Unauthorized")
	}
	r := cache.Data(p.printer)
	r["css"] = cache.CSS
	r["header"] = header(p.printer, p.status)
	r["footer"] = footer(p.printer)
	r["basedir"] = getBaseDir(cache.Path)
	// comparing an int against cache.Extra is not useful, we need an array of keys instead set at AddPlugin.
	for _, key := range pluginKey {
		if (cache.Extra & key) != 0 {
			r[pluginCache[key].Name] = pluginCache[key].Run(p.printer)
		}
	}

	if p.session != nil && cache.Extra&FormErrors != 0 {
		r["errors"] = p.session.Get("errors")
	}

	//if cache.Extra&ClientName != 0 {
	//	i["firstname"] = db.ClientName(p.session)
	//}
	//if cache.Extra&ClientSurname != 0 {
	//	i["surname"] = db.ClientSurname(p.session)
	//}
	//if cache.Extra&ClientUsername != 0 {
	//	i["username"] = db.ClientUsername(p.session)
	//}
	if cache.Extra&FormErrors != 0 && p.session != nil {
		r["errors"] = p.session.Get("errors")
	}

	if p.session != nil {
		r["username"] = p.session.Get("username")
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

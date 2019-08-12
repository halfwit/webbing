package router

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"path"
	"strings"

	"golang.org/x/text/message"
)

var pagecache map[string]*Page
var countrylist []Country

func init() {
	pagecache = make(map[string]*Page)
	countrylist = listcountries()
}

type Access uint8

const (
	GuestAuth Access = 1 << iota
	PatientAuth
	DoctorAuth
)

type IncludeExtra uint8

const (
	ListDoctors IncludeExtra = 1 << iota
	ListServices
	ListCountries
	OneTimeToken
)

type Page struct {
	Access Access
	Extra  IncludeExtra
	Css    string
	Path   string
	Data   func(p *message.Printer) map[string]interface{}
	tmpl   *template.Template
}

func Add(p *Page) {
	pagecache[p.Path+".html"] = p
}

// Walk all our templates and finally return applicable errors as an array
func ValidateAndCache() []error {
	var errs []error
	hd := path.Join("templates", "header.tpl")
	fd := path.Join("templates", "footer.tpl")
	ld := path.Join("templates", "layout.tpl")
	printer := message.NewPrinter(message.MatchLanguage("en"))
	for _, item := range pagecache {
		var err error
		tp := path.Join("templates", item.Path) + ".tpl"

		t := template.New(path.Base(tp))
		item.tmpl, err = t.ParseFiles(tp, hd, fd, ld)
		if err != nil {
			errs = append(errs, fmt.Errorf("parsing in %s - %v", path.Dir(item.Path), err))
			continue
		}
		p := &page{
			printer: printer,
			path: item.Path + ".html",
		}
		_, err = getdata(p, "")
		if err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

func getdata(p *page, in string) ([]byte, error) {
        cache, ok := pagecache[p.path]
	if ! ok {
		return nil,  fmt.Errorf("No such page: %s", p.path)
	}
	i := cache.Data(p.printer)
	i["css"]     = cache.Css
	i["header"]  = header(p.printer, p.status)
	i["footer"]  = footer(p.printer)
	i["basedir"] = getBaseDir(cache.Path)
	if cache.Extra&ListDoctors != 0 {
		i["doctors"] = listdoctors()
	}
	if cache.Extra&ListServices != 0 {
		i["specialties"] = specialties(p.printer)
	}
	if cache.Extra&ListCountries != 0 {
		i["countrylist"] = countrylist
	}
	if p.session != nil && cache.Extra&OneTimeToken != 0 {
		i["token"] = p.session.Get("token")
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
	//if cache.Extra&FormErrors != 0 && p.session != nil {
	//	i["errors"] = p.session.Get("errors")
	//}
	if p.session != nil {
		i["username"] = p.session.Get("username")
		i["errors"] = p.session.Get("errors")
	}
	return cache.render(i)
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

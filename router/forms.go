package router

import (
	"fmt"
	"net/http"

	"github.com/olmaxmedical/session"
	"golang.org/x/text/message"
)

var formlist map[string]*Form

// Form - POST requests
type Form struct {
	Access    Access
	After     PluginMask
	Path      string
	Redirect  string
	Validator func(r *http.Request, p *message.Printer) []string
}

func init() {
	formlist = make(map[string]*Form)
}

// AddPost - Register a POST form from forms/
func AddPost(f *Form) {
	formlist[f.Path+".html"] = f
}

func parseForm(p *Request, w http.ResponseWriter, r *http.Request) (*Form, []string) {
	var errors []string
	form, ok := formlist[p.path]
	if !ok {
		errors = append(errors, "No such page")
		return nil, errors
	}
	if errs := form.Validator(r, p.printer); len(errs) > 0 {
		return nil, errs
	}
	for _, key := range pluginKey {
		if (form.After&key) != 0 && pluginCache[key].Validate != nil {
			if e := pluginCache[key].Validate(p); e != nil {
				errors = append(errors, fmt.Sprint(e))
				return nil, errors
			}
		}
	}
	return form, errors
}

func postform(p *Request, us session.Session, w http.ResponseWriter, r *http.Request) {
	form, errors := parseForm(p, w, r)
	if len(errors) > 0 && errors[0] != "nil" {
		// NOTE(halfwit) this stashes previous entries, but does not work
		// on multipart forms (with file uploads)
		us.Set("errors", errors)
		// Maybe store form args instead here in session
		url := fmt.Sprintf("%s?%s", r.URL.String(), r.Form.Encode())
		http.Redirect(w, r, url, 302)
	}
	if form != nil {
		us.Set("errors", []string{})
		http.Redirect(w, r, form.Redirect, 302)
	}
}

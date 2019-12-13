package router

import (
	"fmt"
	"net/http"

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

func parseform(p *Request, w http.ResponseWriter, r *http.Request) (*Form, []string) {
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

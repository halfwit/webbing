package plugins

import (
	"github.com/olmaxmedical/olmax_go/router"
)

// FormErrors - A list of errors present on a POST request
const FormErrors router.PluginMask = 1 << 7

func init() {
	b := &router.Plugin{
		Name:     "errors",
		Run:      GetFormErrors,
		Validate: nil,
	}
	router.AddPlugin(b, FormErrors)
}

// GetFormErrors - return the client a list of any errors in the form
func GetFormErrors(r *router.Request) map[string]interface{} {
	s := r.Session()
	if s == nil {
		return nil
	}
	return map[string]interface{}{
		"errors": s.Get("errors"),
	}
}

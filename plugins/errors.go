package plugins

import (
	"github.com/olmaxmedical/olmax_go/router"
)

// FormErrors - A list of errors present on a POST request
const FormErrors router.PluginMask = 6

func init() {
	b := &router.Plugin{
		Name:     "Form Errors",
		Run:      GetFormErrors,
		Validate: nil,
	}
	router.AddPlugin(b, FormErrors)
}

// GetFormErrors - Return all errors encountered during form parse
func GetFormErrors(r *router.Request) map[string]interface{} {
	s := r.Session()
	if s == nil {
		return nil
	}
	return map[string]interface{}{
		"errors": s.Get("errors"),
	}
}

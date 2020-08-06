package forms

import (
	"net/http"

	"github.com/albrow/forms"
	"github.com/olmaxmedical/router"
	"golang.org/x/text/message"
)

func init() {
	b := &router.Form{
		Access:    router.PatientAuth,
		Path:      "patient/profile",
		Validator: profile,
		After:     0,
		Redirect:  "/patient/profile.html",
	}
	router.AddPost(b)
}

func profile(r *http.Request, p *message.Printer) []string {
	var errors []string
	data, err := forms.Parse(r)
	if err != nil {
		errors = append(errors, p.Sprint("Internal server error"))
		return errors
	}
	val := data.Validator()
	//
	if val.HasErrors() {
		errors = append(errors, val.Messages()...)
	}
	return errors
}

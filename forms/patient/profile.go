package forms

import (
	"net/http"
	"github.com/albrow/forms"
	"golang.org/x/text/message"
	"olmax/router"
)

func init() {
	b := &router.Form{
		Access: router.PatientAuth,
		Path: "patient/profile",
		Validator: PatientProfile,
		After: 0,
		Redirect: "/patient/profile.html",
	}
	router.AddPost(b)
}

func PatientProfile(r *http.Request, p *message.Printer) []string {
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

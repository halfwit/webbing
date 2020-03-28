package forms

import (
	"net/http"

	"github.com/albrow/forms"
	"github.com/olmaxmedical/plugins"
	"github.com/olmaxmedical/router"
	"golang.org/x/text/message"
)

func init() {
	b := &router.Form{
		Access:    router.GuestAuth,
		Path:      "resetpassword",
		Validator: reset,
		Redirect:  "/login.html",
		After:     plugins.ResetPassword,
	}
	router.AddPost(b)
}

func reset(r *http.Request, p *message.Printer) []string {
	var errors []string
	data, err := forms.Parse(r)
	if err != nil {
		errors = append(errors, "Internal server error")
		return errors
	}
	val := data.Validator()
	val.Require("email").Message(p.Sprintf("Valid email required"))
	val.MatchEmail("email").Message(p.Sprintf("Invalid email"))
	if val.HasErrors() {
		errors = append(errors, val.Messages()...)
	}
	return errors
}

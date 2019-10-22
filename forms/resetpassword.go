package forms

import (
	"net/http"
	"github.com/albrow/forms"
	"golang.org/x/text/message"
	"olmax/router"
)

func init() {
	b := &router.Form{
		Access: router.GuestAuth,
		Path: "resetpassword",
		Validator: Reset,
		Redirect: "/login.html",
		After: router.SendReset,
	}
	router.AddPost(b)
}

func Reset(r *http.Request, p *message.Printer) []string {
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

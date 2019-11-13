package forms

import (
	"net/http"

	"github.com/albrow/forms"
	"github.com/olmaxmedical/olmax_go/plugins"
	"github.com/olmaxmedical/olmax_go/router"
	"golang.org/x/text/message"
)

func init() {
	b := &router.Form{
		Access:    router.GuestAuth,
		Path:      "signup",
		Validator: signin,
		Redirect:  "/login.html",
		After:     plugins.SendSignup,
	}
	router.AddPost(b)
}

func signin(r *http.Request, p *message.Printer) []string {
	var errors []string
	data, err := forms.Parse(r)
	if err != nil {
		errors = append(errors, "Internal server error")
		return errors
	}
	val := data.Validator()
	val.Require("fname").Message(p.Sprintf("First name required"))
	val.MinLength("fname", 2).Message(p.Sprintf("First name must be at least 2 characters"))
	val.Require("lname").Message(p.Sprintf("Last name required"))
	val.MinLength("lname", 2).Message(p.Sprintf("Last name must be at least 2 characters"))
	val.Require("email").Message(p.Sprintf("Valid email required"))
	val.MatchEmail("email").Message(p.Sprintf("Invalid email"))
	val.Require("pass").Message(p.Sprintf("Password required"))
	val.MinLength("pass", 8).Message(p.Sprintf("Password must be at least 8 characters"))
	if val.HasErrors() {
		errors = append(errors, val.Messages()...)
	}
	return errors
}

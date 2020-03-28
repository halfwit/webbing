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
		Path:      "newpassword",
		Validator: newPassword,
		Redirect:  "/login.html",
		After:     plugins.ResetPassword | plugins.FormToken,
	}
	router.AddPost(b)
}

func newPassword(r *http.Request, p *message.Printer) []string {
	var errors []string
	data, err := forms.Parse(r)
	if err != nil {
		errors = append(errors, "Internal server error")
		return errors
	}
	val := data.Validator()
	val.Require("password").Message(p.Sprintf("Password required"))
	val.MinLength("password", 8).Message(p.Sprintf("Password must be at least 8 characters"))
	val.Require("reenter").Message(p.Sprintf("Re-enter same password"))
	val.MinLength("reenter", 8).Message(p.Sprintf("Password must be at least 8 characters"))
	if val.HasErrors() {
		errors = append(errors, val.Messages()...)
	}
	if data.Get("reenter") != data.Get("password") {
		errors = append(errors, p.Sprint("Passwords do not match"))
	}
	return errors
}

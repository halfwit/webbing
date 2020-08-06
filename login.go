package pages

import (
	"github.com/olmaxmedical/plugins"
	"github.com/olmaxmedical/router"
	"golang.org/x/text/message"
)

func init() {
	b := &router.Page{
		Access: router.GuestAuth,
		CSS:    "",
		Path:   "login",
		Data:   login,
		Extra:  plugins.FormErrors,
	}
	router.AddPage(b)
}

func login(p *message.Printer) map[string]interface{} {
	// TODO: Also add in the error messages here
	return map[string]interface{}{
		"title":          p.Sprint("Olmax Medical | Login"),
		"continue":       p.Sprint("Please login to continue"),
		"greeting":       p.Sprint("Welcome back!"),
		"email":          p.Sprint("Email:"),
		"password":       p.Sprint("Password:"),
		"forgotPassword": p.Sprint("Forgot your password?"),
		"login":          p.Sprint("Login"),
	}
}

package pages

import (
	"github.com/olmaxmedical/olmax_go/plugins"
	"github.com/olmaxmedical/olmax_go/router"
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
		"title":          p.Sprintf("Olmax Medical | Login"),
		"greeting":       p.Sprintf("Welcome back!"),
		"email":          p.Sprintf("Email:"),
		"password":       p.Sprintf("Password:"),
		"forgotPassword": p.Sprintf("Forgot your password?"),
		"login":          p.Sprintf("Login"),
	}
}

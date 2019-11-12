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
		Data:   Login,
		Extra:  plugins.FormErrors,
	}
	//router.AddGet(b)
	router.AddPage(b)
}

// Login - olmaxmedical.com/login.html
func Login(p *message.Printer) map[string]interface{} {
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

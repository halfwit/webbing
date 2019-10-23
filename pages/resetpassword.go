package pages

import (
	"github.com/olmaxmedical/olmax_go/router"
	"golang.org/x/text/message"
)

func init() {
	b := &router.Page{
		Access: router.GuestAuth,
		Css:    "",
		Path:   "resetpassword",
		Data:   ResetPassword,
		Extra:  router.FormErrors,
	}
	router.Add(b)
}

func ResetPassword(p *message.Printer) map[string]interface{} {
	// TODO: Also add in the error messages here
	return map[string]interface{}{
		"title":     p.Sprintf("Olmax Medical | Login"),
		"reset":     p.Sprintf("Enter Email"),
		"resettext": p.Sprintf("We will send a reset code to the email provided"),
		"email":     p.Sprintf("Email:"),
		"sendreset": p.Sprintf("Reset"),
	}
}

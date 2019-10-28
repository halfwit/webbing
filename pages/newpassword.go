package pages

import (
	"github.com/olmaxmedical/olmax_go/router"
	"golang.org/x/text/message"
)

func init() {
	b := &router.Page{
		Access: router.GuestAuth,
		CSS:    "",
		Path:   "newpassword",
		Data:   NewPassword,
		Extra:  router.FormToken | router.FormErrors,
	}
	router.Add(b)
}

func NewPassword(p *message.Printer) map[string]interface{} {
	return map[string]interface{}{
		"title":    p.Sprintf("Olmax Medical | Login"),
		"reset":    p.Sprint("Enter new password"),
		"password": p.Sprint("Enter password"),
		"reenter":  p.Sprint("Re-enter password"),
		"update":   p.Sprint("Update"),
	}
}

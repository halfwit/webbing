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
		Path:   "newpassword",
		Data:   newPassword,
		Extra:  plugins.FormToken | plugins.FormErrors,
	}
	router.AddPage(b)
}

func newPassword(p *message.Printer) map[string]interface{} {
	return map[string]interface{}{
		"title":    p.Sprintf("Olmax Medical | Login"),
		"reset":    p.Sprint("Enter new password"),
		"password": p.Sprint("Enter password"),
		"reenter":  p.Sprint("Re-enter password"),
		"update":   p.Sprint("Update"),
	}
}

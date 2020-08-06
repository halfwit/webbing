package pages

import (
	"github.com/olmaxmedical/router"
	"golang.org/x/text/message"
)

func init() {
	b := &router.Page{
		Access: router.PatientAuth | router.DoctorAuth,
		CSS:    "",
		Path:   "messages",
		Data:   messages,
		Extra:  0,
	}
	router.AddPage(b)
}

func messages(p *message.Printer) map[string]interface{} {
	return map[string]interface{}{
		"title":      p.Sprintf("Olmax Medical | Messages"),
		"mainHeader": p.Sprintf("You currently have no messages."),
		"messages":   p.Sprintf("Previous messages: Click here"),
	}
}

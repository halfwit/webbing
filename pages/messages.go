package pages

import (
	"github.com/olmaxmedical/olmax_go/router"
	"golang.org/x/text/message"
)

func init() {
	b := &router.Page{
		Access: router.PatientAuth | router.DoctorAuth,
		CSS:    "",
		Path:   "messages",
		Data:   Messages,
		Extra:  0,
	}
	router.Add(b)
}

// Messages - olmaxmedical.com/messages.html
func Messages(p *message.Printer) map[string]interface{} {
	return map[string]interface{}{
		"title":      p.Sprintf("Olmax Medical | Messages"),
		"mainHeader": p.Sprintf("You currently have no messages."),
		"messages":   p.Sprintf("Previous messages: Click here"),
	}
}

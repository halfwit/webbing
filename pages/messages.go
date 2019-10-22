package pages

import (
	"golang.org/x/text/message"
	"olmax/router"
)

func init() {
	b := &router.Page{
		Access: router.PatientAuth|router.DoctorAuth,
		Css:    "",
		Path:   "messages",
		Data:   Messages,
		Extra:  0,
	}
	router.Add(b)
}

func Messages(p *message.Printer) map[string]interface{} {
	return map[string]interface{}{
		"title":      p.Sprintf("Olmax Medical | Messages"),
		"mainHeader": p.Sprintf("You currently have no messages."),
		"messages":   p.Sprintf("Previous messages: Click here"),
	}
}

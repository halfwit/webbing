package patient

import (
	"github.com/olmaxmedical/router"
	"golang.org/x/text/message"
)

func init() {
	b := &router.Page{
		Access: router.PatientAuth,
		CSS:    "",
		Path:   "patient/profile",
		Data:   profile,
		Extra:  0, // listPendingAppointments
	}
	router.AddPage(b)
}

func profile(p *message.Printer) map[string]interface{} {
	return map[string]interface{}{
		"title":          p.Sprintf("Olmax Medical | Profile"),
		"greetingHeader": p.Sprintf("Hello "),
	}
}

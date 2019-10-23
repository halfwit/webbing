package patient

import (
	"github.com/olmaxmedical/olmax_go/router"
	"golang.org/x/text/message"
)

func init() {
	b := &router.Page{
		Access: router.PatientAuth,
		Css:    "",
		Path:   "patient/profile",
		Data:   Profile,
		Extra:  0, // listPendingAppointments
	}
	router.Add(b)
}

func Profile(p *message.Printer) map[string]interface{} {
	return map[string]interface{}{
		"title":          p.Sprintf("Olmax Medical | Profile"),
		"greetingHeader": p.Sprintf("Hello "),
	}
}

package patient

import (
	"html/template"

	"github.com/olmaxmedical/olmax_go/router"
	"golang.org/x/text/message"
)

func init() {
	b := &router.Page{
		Access: router.PatientAuth,
		CSS:    "",
		Path:   "patient/appointments",
		Data:   Appointments,
		Extra:  0, // call function to look up appointments here
	}
	router.Add(b)
}

// Appointments - olmaxmedical.com/patient/appointments.html
func Appointments(p *message.Printer) map[string]interface{} {
	return map[string]interface{}{
		"title":      p.Sprintf("Olmax Medical | Appointments"),
		"mainHeader": p.Sprintf("You currently have no appointments pending."),
		"mainBody":   p.Sprintf("If you have submitted payment, and do not see appointment scheduled on this page; please refer to the %s section.", template.HTML(`<a href="help.html">help</a>`)),
	}
}

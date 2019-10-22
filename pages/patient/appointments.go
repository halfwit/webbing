package patient

import (
	"html/template"
	"golang.org/x/text/message"
	"olmax/router"
)

func init() {
	b := &router.Page{
		Access: router.PatientAuth,
		Css:    "",
		Path:   "patient/appointments",
		Data:   Appointments,
		Extra:  0,
	}
	router.Add(b)
}

func Appointments(p *message.Printer) map[string]interface{} {
	return map[string]interface{} {
		"title":      p.Sprintf("Olmax Medical | Appointments"),
		"mainHeader": p.Sprintf("You currently have no appointments pending."),
		"mainBody":   p.Sprintf("If you have submitted payment, and do not see appointment scheduled on this page; please refer to the %s section.", template.HTML(`<a href="help.html">help</a>`)),
	}
}

package patient

import (
	"github.com/olmaxmedical/plugins"
	"github.com/olmaxmedical/router"
	"golang.org/x/text/message"
)

func init() {
	b := &router.Page{
		Access: router.PatientAuth,
		CSS:    "",
		Path:   "patient/offer",
		Data:   offer,
		Extra:  plugins.Services | plugins.FormErrors,
	}
	router.AddPage(b)
}

func offer(p *message.Printer) map[string]interface{} {
	return map[string]interface{}{
		"title":           p.Sprint("Olmax Medical | Create Offer"),
		"mainHeader":      p.Sprint("Create An Offer"),
		"specialty":       p.Sprint("Specialty"),
		"specialtyHeader": p.Sprint("Physician Specialty"),
		"bcu":             p.Sprint("Bitcoin Per Unit"),
		"rate":            p.Sprint("15/min"),
		"dates":           p.Sprint("Dates"),
		"from":            p.Sprint("From: "),
		"to":              p.Sprint("To: "),
		"deploy":          p.Sprint("Deploy Contract"),
	}
}

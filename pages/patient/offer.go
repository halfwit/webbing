package patient

import (
	"golang.org/x/text/message"
	"olmax/router"
)

func init() {
	b := &router.Page{
		Access: router.PatientAuth,
		Css:    "",
		Path:   "patient/offer",
		Data:   Createoffer,
		Extra:  router.ListServices|router.FormErrors,
	}
	router.Add(b)
}

func Createoffer(p *message.Printer) map[string]interface{} {
	return map[string]interface{}{
		"title":           p.Sprint("Olmax Medical | Create Offer"),
		"mainHeader":      p.Sprint("Create An Offer"),
		"specialty":	   p.Sprint("Specialty"),
		"specialtyHeader": p.Sprint("Physician Specialty"),
		"bcu":             p.Sprint("Bitcoin Per Unit"),
		"rate":		   p.Sprint("15/min"),
		"dates":           p.Sprint("Dates"),
		"from":            p.Sprint("From: "),
		"to":              p.Sprint("To: "),
		"deploy":          p.Sprint("Deploy Contract"),
	}
}

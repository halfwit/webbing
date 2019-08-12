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
		Extra:  router.ListDoctors,
	}
	router.Add(b)
}

func Createoffer(p *message.Printer) map[string]interface{} {
	return map[string]interface{}{
		"title":           p.Sprintf("Olmax Medical | Create Offer"),
		"mainHeader":      p.Sprintf("Create An Offer"),
		"specialtyHeader": p.Sprintf("Physician Specialty"),
		"bcu":             p.Sprintf("Bitcoin Per Unit <i>(15min)</i>"),
		"dates":           p.Sprintf("Dates"),
		"from":            p.Sprintf("From: "),
		"to":              p.Sprintf("To: "),
		"deploy":          p.Sprintf("Deploy Contract"),
	}
}

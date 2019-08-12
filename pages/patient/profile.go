package patient

import (
	"golang.org/x/text/message"
	"olmax/router"
)

func init() {
	b := &router.Page{
		Access: router.PatientAuth,
		Css:    "",
		Path:   "patient/profile",
		Data:   Profile,
		Extra:  router.ListServices|router.ListCountries,
	}
	router.Add(b)
}

func Profile(p *message.Printer) map[string]interface{} {
	return map[string]interface{}{
		"title":          p.Sprintf("Olmax Medical | Profile"),
		"greetingHeader": p.Sprintf("Hello "),
		"legend":         p.Sprintf("Find A Doctor"),
		"specialty":      p.Sprintf("Specialty"),
		"country":        p.Sprintf("Country"),
		"apptLegend":     p.Sprintf("Appointment Dates: "),
		"from":           p.Sprintf("From:"),
		"to":             p.Sprintf("To:"),
		"search":         p.Sprintf("Search"),
	}
}

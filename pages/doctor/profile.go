package doctor

import (
	"github.com/olmaxmedical/olmax_go/router"
	"golang.org/x/text/message"
)

func init() {
	b := &router.Page{
		Access: router.DoctorAuth,
		CSS:    "",
		Path:   "doctor/profile",
		Data:   Profile,
		Extra:  router.FormErrors | router.FormToken,
	}
	router.AddPage(b)
}

// Profile - olmaxmedical.com/doctor/profile.html
func Profile(p *message.Printer) map[string]interface{} {
	return map[string]interface{}{
		"title":          p.Sprint("Olmax Medical | Profile"),
		"greetingHeader": p.Sprint("Hello "),
		"offer":          p.Sprint("Create New Offer"),
		"specialty":      p.Sprint("Your specialties"),
		"country":        p.Sprint("Your countries"),
		"apptLegend":     p.Sprint("Appointment Times: "),
		"from":           p.Sprint("From:"),
		"to":             p.Sprint("To:"),
		"search":         p.Sprint("Search"),
		"bcu":            p.Sprint("Bitcoin per unit (BTC/15min)"),
		"create":         p.Sprint("Create"),
	}
}

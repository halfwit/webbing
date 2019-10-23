package doctor

import (
	"github.com/olmaxmedical/olmax_go/router"
	"golang.org/x/text/message"
)

func init() {
	b := &router.Page{
		Access: router.DoctorAuth,
		Css:    "",
		Path:   "doctor/bookings",
		Data:   Bookings,
		Extra:  0,
	}
	router.Add(b)
}

func Bookings(p *message.Printer) map[string]interface{} {
	return map[string]interface{}{
		"title":      p.Sprintf("Olmax Medical | Bookings"),
		"mainHeader": p.Sprintf("Available patients"),
		// more fields to populate when we have db access
	}
}

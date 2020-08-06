package doctor

import (
	"github.com/olmaxmedical/router"
	"golang.org/x/text/message"
)

func init() {
	b := &router.Page{
		Access: router.DoctorAuth,
		CSS:    "",
		Path:   "doctor/bookings",
		Data:   bookings,
		Extra:  0,
	}
	router.AddPage(b)
}

func bookings(p *message.Printer) map[string]interface{} {
	//TODO(halfwit) populate with database
	return map[string]interface{}{
		"title":      p.Sprintf("Olmax Medical | Bookings"),
		"mainHeader": p.Sprintf("Available patients"),
	}
}

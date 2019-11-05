package help

import (
	"github.com/olmaxmedical/olmax_go/router"
	"golang.org/x/text/message"
)

func init() {
	b := &router.Page{
		Access: router.GuestAuth | router.PatientAuth | router.DoctorAuth,
		CSS:    "",
		Path:   "help/catalog",
		Data:   Catalog,
		Extra:  0,
	}
	router.Add(b)
}

// Catalog - olmaxmedical.com/help/catalog.html
func Catalog(p *message.Printer) map[string]interface{} {
	return map[string]interface{}{
		"title":      p.Sprintf("Olmax Medical | Our Doctors"),
		"mainHeader": p.Sprintf("Olmax Medical"),
	}
}

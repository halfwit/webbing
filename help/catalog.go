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
		Data:   catalog,
		Extra:  0,
	}
	router.AddPage(b)
}

func catalog(p *message.Printer) map[string]interface{} {
	return map[string]interface{}{
		"title":      p.Sprintf("Olmax Medical | Our Doctors"),
		"mainHeader": p.Sprintf("Olmax Medical"),
	}
}

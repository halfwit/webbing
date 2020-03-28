package doctor

import (
	"github.com/olmaxmedical/olmax_go/router"
	"golang.org/x/text/message"
)

func init() {
	b := &router.Page{
		Access: router.DoctorAuth,
		CSS:    "",
		Path:   "doctor/findpatients",
		Data:   findpatients,
		Extra:  0,
	}
	router.AddPage(b)
}

func findpatients(p *message.Printer) map[string]interface{} {
	return map[string]interface{}{
		"title":      p.Sprintf("Olmax Medical | Find Patients"),
		"mainHeader": p.Sprintf("Available patients"),
		// more fields to populate when we have db access
	}
}

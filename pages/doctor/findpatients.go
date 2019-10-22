package doctor

import (
	"golang.org/x/text/message"
	"olmax/router"
)

func init() {
	b := &router.Page{
		Access: router.DoctorAuth,
		Css:    "",
		Path:   "doctor/findpatients",
		Data:   Findpatients,
		Extra:  0,
	}
	router.Add(b)
}

func Findpatients(p *message.Printer) map[string]interface{} {
	return map[string]interface{}{
		"title":      p.Sprintf("Olmax Medical | Find Patients"),
		"mainHeader": p.Sprintf("Available patients"),
		// more fields to populate when we have db access
	}
}

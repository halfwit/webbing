package doctor

import (
	"golang.org/x/text/message"
	"olmax/router"
)

func init() {
	b := &router.Page{
		Access: router.DoctorAuth,
		Css:    "",
		Path:   "doctor/application",
		Data:   Application,
		Extra:  0,
	}
	router.Add(b)
}

func Application(p *message.Printer) map[string]interface{} {
	return map[string]interface{}{
		"title": p.Sprintf("Olmax Medical | Application"),
	}
}

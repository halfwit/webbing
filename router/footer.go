package router

import (
	"golang.org/x/text/message"
)

func footer(p *message.Printer) map[string]string {
	return map[string]string{
		"faq":      p.Sprintf("FAQ"),
		"help":     p.Sprintf("Help"),
		"banner":   p.Sprintf("Over 12B patients served"),
		"pay":      p.Sprintf("Payment Methods"),
		"fees":     p.Sprintf("Prices and Fees"),
		"verify":   p.Sprintf("Verification"),
		"appt":     p.Sprintf("Appointments"),
		"legal":    p.Sprintf("Legal"),
		"privacy":  p.Sprintf("Privacy Policy"),
		"about":    p.Sprintf("Olmax"),
		"partHead": p.Sprintf("Work with us!"),
		"partner":  p.Sprintf("Become A Partner"),
		"provider": p.Sprint("Become A Provider"),
		"copy":     p.Sprintf("Copyright 2017, 2018, 2019"),
	}
}

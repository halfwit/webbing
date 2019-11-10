package router

import (
	"golang.org/x/text/message"
)

func footer(p *message.Printer) map[string]string {
	return map[string]string{
		"faq":      p.Sprintf("FAQ"),
		"help":     p.Sprintf("Help"),
		"banner":   p.Sprintf("Quality Healthcare"),
		"pay":      p.Sprintf("Payment Methods"),
		"fees":     p.Sprintf("Prices and Fees"),
		"verify":   p.Sprintf("Verification"),
		"appt":     p.Sprintf("Appointments"),
		"legal":    p.Sprintf("Legal"),
		"privacy":  p.Sprintf("Privacy Policy"),
		"howworks": p.Sprint("How It Works"),
		"contact":  p.Sprint("Contact Us"),
		"pricing":  p.Sprint("Pricing"),
		"catalog":  p.Sprint("Catalog"),
		"appts":    p.Sprint("Appointments"),
		"proc":     p.Sprint("Payment Procedures"),
		"payments": p.Sprint("Payment Methods"),
		"phone":    p.Sprint("Call toll free"),
		"number":   p.Sprint("1(555)555-1234"),
		"email":    p.Sprint("Email"),
		"partHead": p.Sprintf("Work with us!"),
		"partner":  p.Sprintf("Become A Partner"),
		"provider": p.Sprint("Become A Provider"),
		"copy":     p.Sprintf("Copyright 2017, 2018, 2019"),
	}
}

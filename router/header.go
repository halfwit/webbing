package router

import (
	"golang.org/x/text/message"
)

func header(p *message.Printer, status string) map[string]string {
	return map[string]string{
		// These go away, in the layout.go they'll be called these values added
		"howworks":  p.Sprint("How It Works"),
		"contact":   p.Sprint("Contact Us"),
		"faq":       p.Sprint("FAQ"),
		"pricing":   p.Sprint("Pricing"),
		"catalog":   p.Sprint("Catalog"),
		"appts":     p.Sprint("Appointments"),
		"proc":      p.Sprint("Payment Procedures"),
		"payments":  p.Sprint("Payment Methods"),
		"fees":      p.Sprint("Prices and Fees"),
		"verify":    p.Sprint("Verification"),
		"phone":     p.Sprint("Call toll free"),
		"number":    p.Sprint("1(555)555-1234"),
		"email":     p.Sprint("Email"),
		"login":     p.Sprint("Login"),
		"logout":    p.Sprint("Logout"),
		"signup":    p.Sprint("Sign Up"),
		"profile":   p.Sprint("Profile"),
		"status":    status,
	}
}

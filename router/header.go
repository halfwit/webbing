package router

import (
	"golang.org/x/text/message"
)

func header(p *message.Printer, status string) map[string]string {
	return map[string]string{
		// These go away, in the layout.go they'll be called these values added
		"doctors":   p.Sprintf("Doctors"),
		"provider":  p.Sprintf("Become A Provider"),
		"whodoctor": p.Sprintf("Who can become a doctor"),
		"howworks":  p.Sprintf("How It Works"),
		"contact":   p.Sprintf("Contact Us"),
		"faq":       p.Sprintf("FAQ"),
		"pricing":   p.Sprintf("Pricing"),
		"appts":     p.Sprintf("Appointments"),
		"proc":      p.Sprintf("Payment Procedures"),
		"payments":  p.Sprintf("Payment Methods"),
		"fees":      p.Sprintf("Prices and Fees"),
		"verify":    p.Sprintf("Verification"),
		"phone":     p.Sprintf("Call toll free"),
		"number":    p.Sprintf("1(555)555-1234"),
		"email":     p.Sprintf("Email"),
		"login":     p.Sprintf("Login"),
		"logout":    p.Sprintf("Logout"),
		"signup":    p.Sprintf("Sign Up"),
		"profile":   p.Sprintf("Profile"),
		"status":    status,
	}
}

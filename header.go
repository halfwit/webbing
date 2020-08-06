package router

import (
	"golang.org/x/text/message"
)

func header(p *message.Printer, status string) map[string]string {
	return map[string]string{
		"home":    p.Sprint("Home"),
		"login":   p.Sprint("Login"),
		"logout":  p.Sprint("Logout"),
		"signup":  p.Sprint("Sign Up"),
		"profile": p.Sprint("Profile"),
		"status":  status,
	}
}

package help

import (
	"html/template"
	"golang.org/x/text/message"
	"olmax/router"
)

func init() {
	b := &router.Page{
		Access: router.GuestAuth,
		Css:    "",
		Path:   "help/verification",
		Data:   Verification,
		Extra:  0,
	}
	router.Add(b)
}

func Verification(p *message.Printer) map[string]interface{} {
	return map[string]interface{}{
		"title":        p.Sprintf("Olmax Medical | FAQ"),
		"mainHeader":   p.Sprintf("Olmax Medical"),
		"verifyHeader": p.Sprintf("Verification"),
		"verifyBody":   p.Sprintf("Profile information such as government ID, diplomas, phone numbers, and emails will will be verified before being posted on Olmax Medical website."),
		"phoneHeader":  p.Sprintf("How do I verify my phone number?"),
		"phoneBody":    p.Sprintf("Once you have submitted a phone number, you can either receive a text message or call with a confirmation number."),
		"noNoteHeader": p.Sprintf("Why did I not get a notification email?"),
		"noNoteBody":   p.Sprintf("Please refer to the %v help page", template.HTML(`<a href="https://olmaxmedical.com/appointmentRequests.html">Appointment Requests</a>`)),
		"blockHeader":  p.Sprintf("If your country blocks Olmax Medical?"),
		"blockBody":    p.Sprintf("You can bypass their firewall using tunnel software such as a VPN or %v software", template.HTML(`<a href="https://www.torproject.org">Tor</a>`)),
	}
}

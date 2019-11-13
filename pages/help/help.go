package help

import (
	"github.com/olmaxmedical/olmax_go/router"
	"golang.org/x/text/message"
)

func init() {
	b := &router.Page{
		Access: router.GuestAuth | router.PatientAuth | router.DoctorAuth,
		CSS:    "",
		Path:   "help/faq",
		Data:   faq,
		Extra:  0,
	}
	router.AddPage(b)
}

func faq(p *message.Printer) map[string]interface{} {
	return map[string]interface{}{
		"title":             p.Sprintf("Olmax Medical | Help"),
		"mainHeader":        p.Sprintf("Olmax Medical"),
		"helpHeader":        p.Sprintf("How can we help?"),
		"topics":            p.Sprintf("Suggested Topics"),
		"appointmentHeader": p.Sprintf("Appointment Requests"),
		"appointmentStatus": p.Sprintf("I am a patient. How do I check the status of my appointment?"),
		"appointmentClear":  p.Sprintf("Should I clear my schedule if I have not heard back from the doctor?"),
		"appointmentExpire": p.Sprintf("What happens if my appointment request is declined or expires?"),
		"paymentHeader":     p.Sprintf("Payment Methods"),
		"paymentEdit":       p.Sprintf("How do I edit or remove a payment method?"),
		"paymentBitcoin":    p.Sprintf("What is Bitcoin?"),
		"paymentBitcoinHow": p.Sprintf("How do I use Bitcoin to pay?"),
		"paymentAdd":        p.Sprintf("How can I add another appointment or business address to my receipt?"),
		"verify":            p.Sprintf("Verification"),
		"verifyPhone":       p.Sprintf("How do I verify my phone number?"),
		"verifyEmail":       p.Sprintf("Why didn't I get my email notification?"),
		"verifyLicense":     p.Sprintf("What is a Verified Medical License?"),
		"priceHeader":       p.Sprintf("Prices & Fees"),
		"priceDetermined":   p.Sprintf("How is the price determined for my appointment?"),
		"priceWhen":         p.Sprintf("When will I be charged?"),
		"priceCurrency":     p.Sprintf("Can I pay with any currency?"),
		"contactHeader":     p.Sprintf("Contacting A Physician"),
		"contactStatus":     p.Sprintf("What does each appointment status mean?"),
		"contactHow":        p.Sprintf("How do I make an appointment on Olmax?"),
		"viewall":           p.Sprintf("-View all"),
	}
}

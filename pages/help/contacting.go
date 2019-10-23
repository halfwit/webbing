package help

import (
	"github.com/olmaxmedical/olmax_go/router"
	"golang.org/x/text/message"
)

func init() {
	b := &router.Page{
		Access: router.GuestAuth | router.PatientAuth | router.DoctorAuth,
		Css:    "",
		Path:   "help/contacting",
		Data:   Contacting,
		Extra:  0,
	}
	router.Add(b)
}

func Contacting(p *message.Printer) map[string]interface{} {
	return map[string]interface{}{
		"title":          p.Sprintf("Olmax Medical | FAQ"),
		"mainHeader":     p.Sprintf("Olmax Medical"),
		"contactHeader":  p.Sprintf("Contacting Physician"),
		"contactBody":    p.Sprintf("You may make contact with your doctor as soon as the appointment is confirmed."),
		"scheduleHeader": p.Sprintf("Do I need to pay before scheduling an appointment?"),
		"scheduleBody":   p.Sprintf("Yes, you must submit payment in order to secure appointment contract. Your payment will be held in escrow until the visit is finalized. Once you submit fees, we will contact the physician and give him or her your medical information. The doctor will then confirm appointment, and an email or text will be sent to you, along with the physicians contact information. Fees are structured in amount of bitcoin (BTC) per unit(U) time (BTC/U). Every unit (U) is equivalent to 15 min, time spent in visit will be pre-determined, and visits going longer that what was agreed upon will not cost extra. All new consults must be a minimum of 2 units, and repeat visits can be a minimum of 1 unit."),
		"chargedHeader":  p.Sprintf("When will I be charged?"),
		"chargedBody":    p.Sprintf("Bitcoin must be paid in full upon deployment or acceptance of contract."),
		"anycurrHeader":  p.Sprintf("Can I pay with any currency?"),
		"anycurrBody":    p.Sprintf("No."),
		"blocksHeader":   p.Sprintf("If your country blocks Olmax Medical?"),
		"blocksBody":     p.Sprintf("You can bypass their firewall using tunnel software such as a VPN, or Tor software. See the following for more information: "),
	}
}

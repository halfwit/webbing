package help

import (
	"github.com/olmaxmedical/router"
	"golang.org/x/text/message"
)

func init() {
	b := &router.Page{
		Access: router.GuestAuth | router.PatientAuth | router.DoctorAuth,
		CSS:    "",
		Path:   "help/pricesandfees",
		Data:   pricefee,
		Extra:  0,
	}
	router.AddPage(b)
}

func pricefee(p *message.Printer) map[string]interface{} {
	return map[string]interface{}{
		"title":           p.Sprintf("Olmax Medical | FAQ"),
		"mainHeader":      p.Sprintf("Olmax Medical"),
		"priceHeader":     p.Sprintf("Prices & Fees"),
		"priceBody":       p.Sprintf("Prices and fees are for the most part determined by doctors and patients."),
		"determineHeader": p.Sprintf("How is the price determined for my appointment?"),
		"determineBody":   p.Sprintf("Prices are set by who deploys the contract (doctor or patient). Fees are structured in a amount of bitcoin (BTC) per unit. In which a unit of time equals 15 mins. All new consults must be a minimum of 2 units, and repeat visits can be 1 unit."),
		"chargedHeader":   p.Sprintf("When will I be charged?"),
		"chargedBody":     p.Sprintf("Bitcoin must be paid in full upon deployment or acceptance of contract."),
		"currencyHeader":  p.Sprintf("Can I pay with any currency?"),
		"currencyBody":    p.Sprintf("No"),
		"blockHeader":     p.Sprintf("If your country blocks Olmax Medical?"),
		"blockBody":       p.Sprintf("You can bypass their firewall using tunnel software such as a VPN or Tor software. See the following for more information: "),
	}
}

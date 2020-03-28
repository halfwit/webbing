package help

import (
	"github.com/olmaxmedical/router"
	"golang.org/x/text/message"
)

func init() {
	b := &router.Page{
		Access: router.GuestAuth | router.PatientAuth | router.DoctorAuth,
		CSS:    "",
		Path:   "help/paymentmethods",
		Data:   paymethod,
		Extra:  0,
	}
	router.AddPage(b)
}

func paymethod(p *message.Printer) map[string]interface{} {
	return map[string]interface{}{
		"title":            p.Sprint("Olmax Medical | FAQ"),
		"mainHeader":       p.Sprint("Olmax Medical"),
		"paymentHeader":    p.Sprint("Payment Methods"),
		"paymentBody":      p.Sprintf("All payments will be done via "),
		"whatBTCHeader":    p.Sprintf("What is Bitcoin? "),
		"whatBTCBody":      p.Sprint("Bitcoin is a form of electronic cash without a central bank or single administrator that can be sent from user to user on the peer-to-peer bitcoin network without the need for intermediaries."),
		"chargedHeader":    p.Sprint("When will I be charged?"),
		"chargedBody":      p.Sprint("Once you decide on which doctor you would like to make an appointment with, you will be asked to deposit bitcoin fund into a dedicated wallet where it will be held in escrow until visit is complete. If visit is cannot be completed, the funds will be deposited back to the original wallet in full."),
		"anyCoinHeader":    p.Sprint("Can I pay with any currency?"),
		"anyCoinBody":      p.Sprint("No"),
		"editWalletHeader": p.Sprint("How do I edit or remove a bitcoin wallet?"),
		"editWalletBody":   p.Sprint("If you choose to pay with a certain wallet address when you make an appointment request, you can edit or remove wallet address from your Olmax Medical account after bitcoin is out of escrow, pending or active consult."),
		"delWalletHeader":  p.Sprint("Removing or adding personal bitcoin wallets"),
		"delWalletBody":    p.Sprint("You can manage the bitcoin wallets in the Account section of your profile:"),
		"delWalletBody1":   p.Sprint("1.) Click Wallet"),
		"delWalletBody2":   p.Sprint("2.) Click Change Wallet Address to add a new wallet to your account"),
		"otherCoinNote":    p.Sprint("Note: At this time, it's only possible to make payments with bitcoin. All other currencies are not supported."),
		"blockedHeader":    p.Sprint("If your country blocks Olmax Medical?"),
		"blockedBody":      p.Sprintf("You can bypass their firewall using tunnel software such as a VPN or Tor software. See the following for more information: "),
	}
}

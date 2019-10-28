package pages

import (
	"github.com/olmaxmedical/olmax_go/router"
	"golang.org/x/text/message"
)

func init() {
	b := &router.Page{
		Access: router.PatientAuth | router.DoctorAuth,
		CSS:    "",
		Path:   "wallet",
		Data:   Wallet,
		Extra:  0,
	}
	router.Add(b)
}

func Wallet(p *message.Printer) map[string]interface{} {
	return map[string]interface{}{
		"title":      p.Sprintf("Olmax Medical | Wallet"),
		"mainHeader": p.Sprintf("Wallet"),
		"funds":      p.Sprintf("0 BTC"),
		"current":    p.Sprintf("NO FUNDS CURRENTLY HELD IN ESCROW"),
		"deposit":    p.Sprintf("Deposit Funds"),
		"onlyHeader": p.Sprintf("Send only Bitcoin (BTC) to this address"),
		"onlyBody":   p.Sprintf("Sending any other digital asset, including Bitcoin Cash (BCH), will result in permanent loss."),
	}
}

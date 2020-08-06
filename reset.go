package email

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/smtp"

	"github.com/olmaxmedical/database"
	"golang.org/x/text/message"
)

// SendReset - Wrapper for resetmail and timeout
// These all need context
func SendReset(ctx context.Context, db *database.Database, user *database.User, email, password string, p *message.Printer) {
	token := db.CreateTempEntry(ctx, user, email, password)
	resetemail(token, []string{email}, p)
}

// NextResetToken - A user has replied to the email about a password
// Give them a new token and invalidate the old one
func NextResetToken(ctx context.Context, db *database.Database, token string) string {
	return db.ForwardToken(ctx, token)
}

func resetemail(token string, to []string, p *message.Printer) {
	var msg bytes.Buffer

	msg.WriteString("From: ")
	msg.WriteString(gmail + "\n")
	msg.WriteString("To: ")
	msg.WriteString(to[0] + "\n")
	msg.WriteString(p.Sprintf("Subject: Olmax Medical - Reset Your Password\n\n"))
	msg.WriteString(p.Sprintf("Please click the following link to reset your password "))
	msg.WriteString(fmt.Sprintf("%s/reset/%s\n", url, token))

	auth := smtp.PlainAuth("", gmail, pw, addr)
	data := msg.Bytes()

	if e := smtp.SendMail(addr+":587", auth, gmail, to, data); e != nil {
		log.Printf("smtp error: %v", e)
	}
}

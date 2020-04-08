package email

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"net/smtp"

	"github.com/olmaxmedical/database"
	"golang.org/x/text/message"
)

// These all need context
// SendSignup - email our prospective clients and create tokens
func SendSignup(ctx context.Context, db *database.Database, user *database.User, email, password string, p *message.Printer) {
	token := db.CreateTempEntry(ctx, user, email, password)
	signupemail(token, []string{email}, p)
}

// ValidateSignupToken - Make sure token is good
func ValidateSignupToken(ctx context.Context, db *database.Database, w http.ResponseWriter, r *http.Request, token string) {
	if e := db.UserFromTemp(ctx, token); e != nil {
		http.Error(w, "Bad Request", 400)
	}

	http.Redirect(w, r, "/login.html", 302)
	return
}

func signupemail(token string, to []string, p *message.Printer) {
	var msg bytes.Buffer

	msg.WriteString("From: ")
	msg.WriteString(gmail + "\n")
	msg.WriteString("To: ")
	msg.WriteString(to[0] + "\n")
	msg.WriteString(p.Sprintf("Subject: Olmax Medical - Verify your new account\n\n"))
	msg.WriteString(p.Sprintf("Please click the following link to finalize your account creation "))
	msg.WriteString(fmt.Sprintf("%s/activate/%s\n", url, token))

	auth := smtp.PlainAuth("", gmail, pw, addr)
	data := msg.Bytes()

	if e := smtp.SendMail(addr+":587", auth, gmail, to, data); e != nil {
		log.Printf("smtp error: %v", e)
	}
}

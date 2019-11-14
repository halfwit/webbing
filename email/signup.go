package email

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"time"

	"github.com/google/uuid"
	"github.com/olmaxmedical/olmax_go/db"
	"golang.org/x/text/message"
)

// TODO(halfwit) Update with whichever TLD we land on
var url = "https://olmaxmedical.com"

// SendSignup - email our prospective clients and create tokens
func SendSignup(first, last, email, pass string, p *message.Printer) {
	if !db.UserExists(email) {
		u, _ := uuid.NewRandom()
		token := u.String()
		db.CreateTempEntry(first, last, email, pass, token)
		signupemail(token, email, p)
		go func() {
			// Blow away the entry unconditionally after 10 minutes
			time.Sleep(time.Minute * 10)
			db.RemoveTempEntry(token)
		}()
	}
}

// ValidateSignupToken - Make sure token is good
func ValidateSignupToken(w http.ResponseWriter, r *http.Request, token string) {
	if db.FindTempEntry(token) {
		db.CreateEntry(token)
		http.Redirect(w, r, "/login.html", 302)
		return
	}
	http.Error(w, "Bad Request", 400)

}

func signupemail(token string, sendto string, p *message.Printer) {
	var msg bytes.Buffer
	msg.WriteString("From: ")
	msg.WriteString("olmaxmedical@gmail.com" + "\n")
	msg.WriteString("To: ")
	msg.WriteString(sendto + "\n")
	msg.WriteString(p.Sprintf("Subject: Olmax Medical - Verify your new account\n\n"))
	msg.WriteString(p.Sprintf("Please click the following link to finalize your account creation "))
	msg.WriteString(fmt.Sprintf("%s/activate/%s\n", url, token))
	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", "olmaxmedical@gmail.com", "hunter2", "smtp.gmail.com"),
		"olmaxmedical@gmail.com", []string{sendto}, msg.Bytes(),
	)
	if err != nil {
		log.Printf("smtp error: %v", err)
	}
}

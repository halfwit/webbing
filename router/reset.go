package router

import (
	"bytes"
	"fmt"
	"log"
	"net/smtp"
	"time"

	"github.com/google/uuid"
	"golang.org/x/text/message"
	"olmax/db"
)

func nextResetToken(old, user string) string {
	if db.FindTempEntry(old) {
		db.RemoveTempEntry(old)
		u, _ := uuid.NewRandom()
		token := u.String()
		db.CreateTempEntry("", "", user, "", token)
		go func() {
			time.Sleep(time.Minute * 10)
			db.RemoveTempEntry(token)
		}()
		return token
	}
	return ""
}

func sendreset(email string, p *message.Printer) string {
	u, _ := uuid.NewRandom()
	token := u.String()
	if  db.UserExists(email) {	
		db.CreateTempEntry("", "", email, "", token)
		resetemail(token, email, p)
		go func() {
			time.Sleep(time.Minute * 10)
			db.RemoveTempEntry(token)
		}()
	}
	return "<a href=\"https://mail.google.com/mail/u?authuser=" + email + "\"/>"
}

func resetemail(token string, sendto string, p *message.Printer) {
	var msg bytes.Buffer
	msg.WriteString("From: ")
	msg.WriteString("olmaxmedical@gmail.com" + "\n")
	msg.WriteString("To: ")
	msg.WriteString(sendto + "\n")
	msg.WriteString(p.Sprintf("Subject: Olmax Medical - Reset Your Password\n\n"))
	msg.WriteString(p.Sprintf("Please click the following link to reset your password "))
	msg.WriteString(fmt.Sprintf("%s/reset/%s\n", url, token))
	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", "olmaxmedical@gmail.com", "hunter2", "smtp.gmail.com"),
		"olmaxmedical@gmail.com", []string{sendto}, msg.Bytes(),
	)
	if err != nil {
		log.Printf("smtp error: %v", err)
	}
}

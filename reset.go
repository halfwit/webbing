package email

import (
	"bytes"
	"fmt"
	"log"
	"net/smtp"
	"time"

	"github.com/google/uuid"
	"github.com/olmaxmedical/database"
	"golang.org/x/text/message"
)

// SendReset - Wrapper for resetmail and timeout
func SendReset(email string, p *message.Printer) {
	u, _ := uuid.NewRandom()
	token := u.String()
	if database.UserExists(email) {
		database.CreateTempEntry("", "", email, "", token)
		resetemail(token, email, p)
		go func() {
			time.Sleep(time.Minute * 10)
			database.RemoveTempEntry(token)
		}()
	}
}

// NextResetToken - Make sure we have unique tokens!
func NextResetToken(old, user string) string {
	if database.FindTempEntry(old) {
		database.RemoveTempEntry(old)
		u, _ := uuid.NewRandom()
		token := u.String()
		database.CreateTempEntry("", "", user, "", token)
		go func() {
			time.Sleep(time.Minute * 10)
			database.RemoveTempEntry(token)
		}()
		return token
	}
	return ""
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

package plugins

import (
	"errors"
	"log"
	"mime/multipart"

	"github.com/olmaxmedical/olmax_go/email"
	"github.com/olmaxmedical/olmax_go/router"
)

// EmailForm - Patient form to gmail
const EmailForm router.PluginMask = 4

// SendSignup - Send account creation validation email
const SendSignup router.PluginMask = 5

// SendReset - Send password reset email
const SendReset router.PluginMask = 6

func init() {
	b := &router.Plugin{
		Name:     "Email client information form",
		Run:      nil,
		Validate: emailForm,
	}
	router.AddPlugin(b, EmailForm)
	c := &router.Plugin{
		Name:     "Send signup email",
		Run:      nil,
		Validate: signupEmail,
	}
	router.AddPlugin(c, SendSignup)
	d := &router.Plugin{
		Name:     "Send password reset email",
		Run:      nil,
		Validate: resetPassword,
	}
	router.AddPlugin(d, SendReset)
}

func signupEmail(s *router.Request) error {
	r := s.Request()
	first := r.PostFormValue("fname")
	last := r.PostFormValue("lname")
	address := r.PostFormValue("email")
	pass := r.PostFormValue("pass")
	email.SendSignup(first, last, address, pass, s.Printer())
	return nil
}

func resetPassword(s *router.Request) error {
	p := s.Printer()
	r := s.Request()
	email.SendReset(r.PostFormValue("email"), p)
	return nil
}

func emailForm(s *router.Request) error {
	p := s.Printer()
	r := s.Request()
	r.ParseMultipartForm(10 << 20) // parse up to 10MB
	if r.PostFormValue("name") == "" || r.PostFormValue("email") == "" {
		return errors.New(p.Sprint("Missing name or email in form. Please contact us at olmaxmedical@gmail.com"))
	}
	if b, ok := r.Form["sendto"]; !ok || b[0] == "" {
		return errors.New(p.Sprint("Missing value for target email. Please contact us at olmaxmedical.gmail.com"))
	}
	attachments := make(map[string]multipart.File)
	m := r.MultipartForm
	for _, headers := range m.File {
		for _, header := range headers {
			file, err := header.Open()
			if err != nil { //non fatal, log any oddities and continue
				log.Println(err)
				continue
			}
			attachments[header.Filename] = file
		}
	}
	email.SendForm(r.Form, p, attachments)
	return nil
}

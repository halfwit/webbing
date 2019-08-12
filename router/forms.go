package router

import (
	"fmt"
	"net/http"
	"golang.org/x/text/message"
	"olmax/db"
	"olmax/session"
)

var formlist map[string]*Form

type After uint8

const (
	ValidateLogin After = 1 << iota
	SendSignup
	SendReset
	SetPassword
)

type Form struct {
	Access    Access
	After     After
	Path      string
	Redirect  string
	Validator func(r *http.Request, p *message.Printer) []string
}

func init() {
	formlist = make(map[string]*Form)
}

func AddPost(f *Form) {
	formlist[f.Path + ".html"] = f
}

func parseform(p *page, w http.ResponseWriter, r *http.Request) []string {
	var errors []string
	form, ok := formlist[p.path]
	if ! ok {
		errors = append(errors, "No such page")
		return errors
	}
	e := form.Validator(r, p.printer)
	if len(e) > 0 {
		errors = append(errors, e...)
	}
	var msg string
	if form.After&ValidateLogin != 0 {
		e = validateLogin(p.printer, p.session, r)
		if len(e) > 0 {
			errors = append(errors, e...)
		}	
	}
	if form.After&SendSignup != 0 {
		msg = signupEmail(p.printer, r)
	}
	if form.After&SendReset != 0 {
		msg = resetPassword(p.printer, r)
	}
	if form.After&SetPassword != 0 {
		e = setPassword(p.printer, p.session, r)
		if len(e) > 0 {
			errors = append(errors, e...)
		}
	}
	if msg != "" {
		fmt.Fprintf(w, "%s\n", msg)
		return errors
	}
	if len(errors) == 0 {
		http.Redirect(w, r, form.Redirect, 302)
	}
	return errors
}

func setPassword(p *message.Printer, us session.Session, r *http.Request) []string {
	var errors []string
	pass := r.PostFormValue("password")
	repeat := r.PostFormValue("reenter")
	if pass != repeat {
		errors = append(errors, p.Sprint("Passwords do not match"))
		return errors
	}
	token := r.PostFormValue("token")
	if ! db.FindEntry(token) {
		errors = append(errors, p.Sprint("Session expired"))
		return errors
	}
	db.UpdateUserPassword(token, pass)
	return errors
}

func validateLogin(p *message.Printer, us session.Session, r *http.Request) []string {
	var errors []string
	user := r.PostFormValue("email")
	pass := r.PostFormValue("pass")
	if db.ValidateLogin(user, pass) {
		us.Set("username", user)
		us.Set("login", "true")
		return errors
	}
	errors = append(errors, p.Sprint("Invalid username or password"))
	return errors
}

func signupEmail(p *message.Printer, r *http.Request) string {
	first := r.PostFormValue("fname")
	last := r.PostFormValue("lname")
	email := r.PostFormValue("email")
	pass := r.PostFormValue("pass")
	em := sendsignup(first, last, email, pass, p)
	return p.Sprintf("An email has been sent to the following email with instructions on finalizing your account creation: %s\n", em)
}

func resetPassword(p *message.Printer, r *http.Request) string{
	em := sendreset(r.PostFormValue("email"), p)
	return p.Sprintf("An email has been sent to the following email with a link to reset your password: %s\n", em)
}



package router

import (
	"fmt"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/olmaxmedical/olmax_go/db"
	"github.com/olmaxmedical/olmax_go/email"
	"github.com/olmaxmedical/olmax_go/session"
	"golang.org/x/text/message"
)

var formlist map[string]*Form

// After will go away when with plugins
type After uint16

const (
	ValidateLogin After = 1 << iota
	ValidateCountry
	ValidateSpecialty
	ValidateCountries
	ValidateSpecialties
	ValidateToken
	WithOffer
	Search
	SendSignup
	SendReset
	SetPassword
	EmailForm
	AddAppointment
)

// Form - POST requests
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

// AddPost - Register a POST form from forms/
func AddPost(f *Form) {
	formlist[f.Path+".html"] = f
}

// This large ladder just adds conditional logic to forms for a more generic
// Ideally, the *page abstraction will never leak into the form validation
func parseform(p *request, w http.ResponseWriter, r *http.Request) (*Form, []string) {
	var errors, errs []string
	var msg string
	form, ok := formlist[p.path]
	if !ok {
		errors = append(errors, "No such page")
		return nil, errors
	}
	if form.After&ValidateToken != 0 {
		t := r.PostFormValue("token")
		if !validateToken(t) {
			return nil, []string{p.printer.Sprint("Invalid form token")}
		}
	}
	if form.After&WithOffer != 0 {
		//token := r.PostFormValue("sessiontoken")
		//offer := db.GetOffer(token)
		//r.Form["sendto"] = []string{offer.Email}
		//r.Form["offerid"] = []string{offer.Id}
		r.Form["sendto"] = []string{
			"michaelmisch1985@gmail.com",
		}
		r.Form["offerid"] = []string{
			"void",
		}
	}
	if errs = form.Validator(r, p.printer); len(errs) > 0 {
		return nil, errs
	}
	if form.After&ValidateLogin != 0 {
		if errs = validateLogin(p.printer, p.session, r); len(errs) > 0 {
			return nil, errs
		}
	}
	/*
		if form.After&ValidateCountry != 0 {
			c := r.PostFormValue("country")
			if e = validateCountry(p.printer, c); e != "" {
				errors = append(errors, e)
			}
		}
		if form.After&ValidateSpecialty != 0 {
			s := r.PostFormValue("specialty")
			if e = validateSpecialty(p.printer, s); e != "" {
				errors = append(errors, e)
			}
		}
		if form.After&ValidateCountries != 0 {
			c := r.Form["country"]
			if e = validateCountries(p.printer, c); e != "" {
				errors = append(errors, e)
			}
		}
		if form.After&ValidateSpecialties != 0 {
			s := r.Form["specialty"]
			if e = validateSpecialties(p.printer, s); e != "" {
				errors = append(errors, e)
			}
		}*/
	if form.After&SetPassword != 0 {
		if errs = setPassword(p.printer, p.session, r); len(errs) > 0 {
			errors = append(errors, errs...)
		}
	}
	if len(errors) > 0 {
		return nil, errors
	}
	if form.After&SendSignup != 0 {
		msg = signupEmail(p.printer, r)
	}
	if form.After&SendReset != 0 {
		msg = resetPassword(p.printer, r)
	}
	if form.After&EmailForm != 0 {
		msg = emailForm(p.printer, r)
	}
	/* TODO(halfwit) once database is live
	if form.After&AddAppointment != 0 {
		e = db.AddAppointment(r.Form)
	}
	if form.After&Search != 0 {
		results = db.Search(r.Form)
	}
	*/
	if msg != "" {
		fmt.Fprintf(w, "%s\n", msg)
		return nil, []string{"nil"}
	}
	return form, errors
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
	if !db.FindTempEntry(token) {
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
		us.Set("role", db.UserRole(user))
		return errors
	}
	errors = append(errors, p.Sprint("Invalid username or password"))
	return errors
}

func signupEmail(p *message.Printer, r *http.Request) string {
	first := r.PostFormValue("fname")
	last := r.PostFormValue("lname")
	address := r.PostFormValue("email")
	pass := r.PostFormValue("pass")
	email.SendSignup(first, last, address, pass, p)
	return p.Sprint("An email has been sent to the provided email with instructions on finalizing your account creation")
}

func resetPassword(p *message.Printer, r *http.Request) string {
	email.SendReset(r.PostFormValue("email"), p)
	return p.Sprint("An email has been sent to the provided email with a link to reset your password")
}

func emailForm(p *message.Printer, r *http.Request) string {
	r.ParseMultipartForm(10 << 20) // parse up to 10MB
	if r.PostFormValue("name") == "" || r.PostFormValue("email") == "" {
		return p.Sprint("Missing name or email in form. Please contact us at olmaxmedical@gmail.com")
	}
	if b, ok := r.Form["sendto"]; !ok || b[0] == "" {
		return p.Sprint("Missing value for target email. Please contact us at olmaxmedical.gmail.com")
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
	return p.Sprint("Your form has been submitted via email ")
}

package router

import (
	"fmt"
	"net/http"

	"github.com/olmaxmedical/database"
	"github.com/olmaxmedical/email"
	"github.com/olmaxmedical/session"
	"golang.org/x/text/message"
)

// Handle specific endpoints
type handler struct {
	manager *session.Manager
}

func (d *handler) logout(w http.ResponseWriter, r *http.Request) {
	d.manager.Destroy(w, r)
	w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	http.Redirect(w, r, "/index.html", 302)
}

func (d *handler) normal(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	if r.URL.Path == "/" {
		http.Redirect(w, r, "/index.html", 302)
		return
	}
	user, status, us, role := d.getUser(w, r)
	p := &Request{
		printer: userLang(r),
		status:  status,
		request: r,
		user:    user,
		role:    role,
		session: us,
		path:    r.URL.Path[1:],
	}

	switch r.Method {
	case "GET":
		getpage(p, w)
	case "POST":
		postform(p, us, w, r)
	}
}

// TODO: This will require actual client data from the database to populate the page
func (d *handler) profile(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	user, status, us, role := d.getUser(w, r)
	if status == "false" {
		http.Redirect(w, r, "/login.html", 302)
		return
	}
	if rd, ok := us.Get("redirect").(string); ok {
		us.Delete("redirect")
		http.Redirect(w, r, "/"+rd, 302)
		return
	}
	p := &Request{
		printer: userLang(r),
		status:  status,
		session: us,
		user:    user,
		role:    role,
	}
	var data []byte
	var err error
	switch database.UserRole(user) {
	case database.DoctorAuth:
		if role != database.DoctorAuth {
			http.Error(w, "Unauthorized", 401)
			return
		}
		p.path = "doctor/profile.html"
		data, err = getData(p, "doctor")
	case database.PatientAuth:
		if role != database.PatientAuth {
			http.Error(w, "Unauthorized", 401)
			return
		}
		p.path = "patient/profile.html"
		data, err = getData(p, "patient")
	default:
		http.Error(w, "Forbidden", 403)
		return
	}
	if err != nil {
		http.Error(w, "Service Unavailable", 503)
		return
	}
	fmt.Fprintf(w, "%s", data)
}

func (d *handler) activate(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	if len(r.URL.Path) != 46 && r.URL.Path[:9] != "/activate" {
		http.Error(w, "Bad Request", 400)
		return
	}
	email.ValidateSignupToken(w, r, r.URL.Path[10:])
}

func (d *handler) reset(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Path) != 43 && r.URL.Path[:6] != "/reset" {
		http.Error(w, "Bad Request", 400)
		return
	}
	p := userLang(r)
	user, _, us, _ := d.getUser(w, r)
	token := email.NextResetToken(r.URL.Path[7:], user)
	if token == "" {
		us.Set("errors", [1]string{p.Sprint("Token expired")})
		return
	}
	us.Set("token", token)
	r.URL.Path = "/newpassword.html"
	d.normal(w, r)
}

func (d *handler) getUser(w http.ResponseWriter, r *http.Request) (string, string, session.Session, database.Access) {
	us := d.manager.Start(w, r)
	user, ok1 := us.Get("username").(string)
	status, ok2 := us.Get("login").(string)
	role, ok3 := us.Get("role").(database.Access)
	if !ok1 || !ok2 || status != "true" {
		status = "false"
	}
	if !ok3 {
		role = database.GuestAuth
	}
	if status == "true" {
		us.Set("token", database.NewToken())
	}
	return user, status, us, role
}

func userLang(r *http.Request) *message.Printer {
	accept := r.Header.Get("Accept-Language")
	lang := r.FormValue("lang")
	tag := message.MatchLanguage(lang, accept)
	return message.NewPrinter(tag)
}

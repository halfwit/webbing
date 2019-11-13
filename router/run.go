package router

import (
	"fmt"
	"net/http"

	"github.com/olmaxmedical/olmax_go/db"
	"github.com/olmaxmedical/olmax_go/email"
	"github.com/olmaxmedical/olmax_go/session"
	"golang.org/x/text/message"
)

type handle struct {
	manager *session.Manager
}

// Route - All requests pass through here first
func Route(manager *session.Manager) error {
	d := &handle{
		manager: manager,
	}
	css := http.FileServer(http.Dir("resources/css/"))
	jss := http.FileServer(http.Dir("resources/scripts"))
	img := http.FileServer(http.Dir("resources/images/"))
	mux := http.NewServeMux()
	mux.Handle("/css/", http.StripPrefix("/css/", css))
	mux.Handle("/scripts/", http.StripPrefix("/scripts/", jss))
	mux.Handle("/images/", http.StripPrefix("/images/", img))
	mux.HandleFunc("/activate/", d.activate)
	mux.HandleFunc("/reset/", d.reset)
	mux.HandleFunc("/logout.html", d.logout)
	mux.HandleFunc("/profile.html", d.profile)
	mux.HandleFunc("/", d.normal)
	return http.ListenAndServe(":8080", mux)
}

func (d *handle) activate(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Path) != 46 && r.URL.Path[:9] != "/activate" {
		http.Error(w, "Bad Request", 400)
		return
	}
	email.ValidateSignupToken(w, r, r.URL.Path[10:])
}

func (d *handle) reset(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Path) != 43 && r.URL.Path[:6] != "/reset" {
		http.Error(w, "Bad Request", 400)
		return
	}
	p := userLang(r)
	user, _, us, _ := getUser(d, w, r)
	token := email.NextResetToken(r.URL.Path[7:], user)
	fmt.Println(r.URL.Path[7:], token)
	if token == "" {
		us.Set("errors", [1]string{p.Sprint("Token expired")})
		return
	}
	us.Set("token", token)
	r.URL.Path = "/newpassword.html"
	d.normal(w, r)
}

// Request represents an incoming GET/POST
type Request struct {
	printer *message.Printer
	session session.Session
	request *http.Request
	user    string
	status  string
	path    string
	role    db.Access
}

// Printer - returns the client's localized printer handler
func (r *Request) Printer() *message.Printer {
	return r.printer
}

// Session - returns the client's session
func (r *Request) Session() session.Session {
	return r.session
}

// Request - underlying http.Request for forms and such
func (r *Request) Request() *http.Request {
	return r.request
}

func (d *handle) normal(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.Redirect(w, r, "/index.html", 302)
		return
	}
	user, status, us, role := getUser(d, w, r)
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
		get(p, w)
	case "POST":
		post(p, us, w, r)
	}
}

func (d *handle) logout(w http.ResponseWriter, r *http.Request) {
	d.manager.Destroy(w, r)
	http.Redirect(w, r, "/index.html", 302)

}

func post(p *Request, us session.Session, w http.ResponseWriter, r *http.Request) {
	form, errors := parseform(p, w, r)
	if len(errors) > 0 && errors[0] != "nil" {
		// NOTE(halfwit) this stashes previous entries, but does not work
		// on multipart forms (with file uploads)
		us.Set("errors", errors)
		url := fmt.Sprintf("%s?%s", r.URL.String(), r.Form.Encode())
		http.Redirect(w, r, url, 302)
	}
	if form != nil {
		us.Set("errors", []string{})
		http.Redirect(w, r, form.Redirect, 302)
	}
}

func get(p *Request, w http.ResponseWriter) {
	var data []byte
	var err error
	switch db.UserRole(p.user) {
	case db.DoctorAuth:
		data, err = getdata(p, "doctor")
	case db.PatientAuth:
		data, err = getdata(p, "patient")
	default:
		data, err = getdata(p, "guest")
	}
	if err != nil {
		http.Error(w, "Service Unavailable", 503)
		return
	}
	fmt.Fprintf(w, "%s", data)
}

func userLang(r *http.Request) *message.Printer {
	accept := r.Header.Get("Accept-Language")
	lang := r.FormValue("lang")
	tag := message.MatchLanguage(lang, accept)
	return message.NewPrinter(tag)
}

func getUser(d *handle, w http.ResponseWriter, r *http.Request) (string, string, session.Session, db.Access) {
	us := d.manager.Start(w, r)
	user, ok1 := us.Get("username").(string)
	status, ok2 := us.Get("login").(string)
	role, ok3 := us.Get("role").(db.Access)
	if !ok1 || !ok2 || status != "true" {
		status = "false"
	}
	if !ok3 {
		role = db.GuestAuth
	}
	return user, status, us, role
}

// TODO: This will require actual client data from the database to populate the page
func (d *handle) profile(w http.ResponseWriter, r *http.Request) {
	user, status, us, role := getUser(d, w, r)
	if status == "false" {
		http.Error(w, "Unauthorized", 401)
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
	switch db.UserRole(user) {
	case db.DoctorAuth:
		if role != db.DoctorAuth {
			http.Error(w, "Unauthorized", 401)
			return
		}
		p.path = "doctor/profile.html"
		data, err = getdata(p, "doctor")
	case db.PatientAuth:
		if role != db.PatientAuth {
			http.Error(w, "Unauthorized", 401)
			return
		}
		p.path = "patient/profile.html"
		data, err = getdata(p, "patient")
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

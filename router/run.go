package router

import (
	"fmt"
	"net/http"

	"golang.org/x/text/message"
	"olmax/db"
	"olmax/session"
)

type handle struct {
	manager *session.Manager
}

func Route(manager *session.Manager) error {
	d := &handle {
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
	return http.ListenAndServe(":80", mux)
}

func (d *handle) activate(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Path) != 46 && r.URL.Path[:9] != "/activate" {
		http.Error(w, "Bad Request", 400)
		return
	}
	validateSignupToken(w, r, r.URL.Path[10:])
}

func (d *handle) reset(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Path) != 43 && r.URL.Path[:6] != "/reset" {
		http.Error(w, "Bad Request", 400)
		return
	}
	p := userLang(r)
	user, _, us := getUser(d, w, r)
	token := nextResetToken(r.URL.Path[7:], user)
	if token == "" {
		us.Set("errors", [1]string{p.Sprint("Token expired")})
		return
	}
	us.Set("token", token)
	r.URL.Path = "/newpassword.html"
	d.normal(w, r)
}

type page struct {
	printer *message.Printer
	session session.Session
	user    string
	status  string
	path    string
}

func (d *handle) normal(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.Redirect(w, r, "/index.html", 302)
		return
	}
	user, status, us := getUser(d, w, r)
	p := &page {
		printer: userLang(r),
		status: status,
		user: user,
		session: us,
		path: r.URL.Path[1:],
	}
	switch r.Method {
	case "GET":
		get(p, w)
	case "POST":
		post(p, us, w, r);
	}
}

func (d *handle) logout(w http.ResponseWriter, r *http.Request) {
	d.manager.SessionDestroy(w, r)
	http.Redirect(w, r, "/index.html", 302)
	
}

func post(p *page, us session.Session, w http.ResponseWriter, r *http.Request) {
	errors := parseform(p, w, r)
	if len(errors) > 0 {
		us.Set("errors", errors)
		get(p, w)
	}
}

func get(p *page, w http.ResponseWriter) {
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

func getUser(d *handle, w http.ResponseWriter, r *http.Request) (string, string, session.Session) {
	us := d.manager.SessionStart(w, r)
	user, ok1 := us.Get("username").(string)
	status, ok2 := us.Get("login").(string)
	if ! ok1 || ! ok2 || status != "true" {
		status = "false"
	}
	return user, status, us
}

// TODO: This will require actual client data from the database to populate the page
func (d *handle) profile(w http.ResponseWriter, r *http.Request) {
	user, status, us := getUser(d, w, r)
	if status == "false" {
		http.Error(w, "Unauthorized", 401)
		return
	}
	p := &page {
		printer: userLang(r),
		status: status,
		session: us,
		user: user,
	}
	var data []byte
	var err error
	switch db.UserRole(user) {
	case db.DoctorAuth:
		p.path = "doctor/profile.html"
		data, err = getdata(p, "doctor")
	case db.PatientAuth:
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


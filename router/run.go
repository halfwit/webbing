package router

import (
	"crypto/tls"
	"net/http"

	"github.com/olmaxmedical/olmax_go/db"
	"github.com/olmaxmedical/olmax_go/session"
	"golang.org/x/text/message"
)

// Route - All requests pass through here first
func Route(manager *session.Manager) error {
	d := &handler{
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
	//from https://github.com/denji/golang-tls (creative commons)
	srv := &http.Server{
		Addr:         ":8443",
		Handler:      mux,
		TLSConfig:    getTlsConfig(),
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}
	return srv.ListenAndServeTLS("cert.pem", "key.pem")
}

// Some utility functions that are shared across pages and forms
func userLang(r *http.Request) *message.Printer {
	accept := r.Header.Get("Accept-Language")
	lang := r.FormValue("lang")
	tag := message.MatchLanguage(lang, accept)
	return message.NewPrinter(tag)
}

func getUser(d *handler, w http.ResponseWriter, r *http.Request) (string, string, session.Session, db.Access) {
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
	if status == "true" {
		us.Set("token", db.NewToken())
	}
	return user, status, us, role
}

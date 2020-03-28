package router

import (
	"crypto/tls"
	"net/http"

	"github.com/olmaxmedical/olmax_go/session"
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

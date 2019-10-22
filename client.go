package main

import (
	"net/http"
	"log"
	_ "olmax/forms"
	_ "olmax/forms/doctor"
	_ "olmax/forms/patient"
	_ "olmax/pages"
	_ "olmax/pages/doctor"
	_ "olmax/pages/help"
	_ "olmax/pages/patient"
	"olmax/router"
	"olmax/session"
)

//go:generate gotext -srclang=en-US update -out=catalog.go -lang=en-US

func main() {
	sessions, err := session.NewManager("default", "sessions", 3600)
	if err != nil {
		log.Fatalf("Unable to initialize manager %v", err)
	}
	// This is session timeouts, I didn't write this logic and it's very, very broken
	//go sessions.GC()

	errs := router.ValidateAndCache()
	if len(errs) > 0 {
		for _, err := range errs {
			log.Print(err)
		}
		log.Fatal("Unable to continue due to template errors")
	}
	go func() {
		http.ListenAndServe(":6060", nil)
	}()
	log.Fatal(router.Route(sessions))
}

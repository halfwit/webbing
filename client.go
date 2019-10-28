package main

import (
	"log"
	"net/http"

	_ "github.com/olmaxmedical/olmax_go/forms"
	_ "github.com/olmaxmedical/olmax_go/forms/doctor"
	_ "github.com/olmaxmedical/olmax_go/forms/patient"
	_ "github.com/olmaxmedical/olmax_go/pages"
	_ "github.com/olmaxmedical/olmax_go/pages/doctor"
	_ "github.com/olmaxmedical/olmax_go/pages/help"
	_ "github.com/olmaxmedical/olmax_go/pages/patient"
	"github.com/olmaxmedical/olmax_go/router"
	//"github.com/olmaxmedical/olmax_go/plugins"
	"github.com/olmaxmedical/olmax_go/session"
)

//go:generate gotext -srclang=en-US update -out=catalog.go -lang=en-US

func main() {
	sessions, err := session.NewManager("default", "sessions", 3600)
	if err != nil {
		log.Fatalf("Unable to initialize manager %v", err)
	}
	//BUG(halfwit) This is session timeouts, I didn't write this logic and it's very, very broken
	//go sessions.GC()

	errs := router.ValidatePages()
	if len(errs) > 0 {
		for _, err := range errs {
			log.Print(err)
		}
		log.Fatal("Unable to continue due to template errors")
	}
	errs := router.ValidatePlugins()
	go func() {
		http.ListenAndServe(":6060", nil)
	}()
	log.Fatal(router.Route(sessions))
}

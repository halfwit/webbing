package main

import (
	"log"
	// Call our init functions to add our items
	_ "github.com/olmaxmedical/forms"
	_ "github.com/olmaxmedical/forms/doctor"
	_ "github.com/olmaxmedical/forms/patient"
	_ "github.com/olmaxmedical/pages"
	_ "github.com/olmaxmedical/pages/doctor"
	_ "github.com/olmaxmedical/pages/help"
	_ "github.com/olmaxmedical/pages/patient"
	_ "github.com/olmaxmedical/plugins"

	"github.com/olmaxmedical/router"
	"github.com/olmaxmedical/session"
)

//go:generate gotext -srclang=en-US update -out=catalog.go -lang=en-US
// use a working dir instead, passed in as an argument instead

func main() {
	sessions, err := session.NewManager("default", "sessions", 360)
	if err != nil {
		log.Fatalf("Unable to initialize manager %v", err)
	}

	go sessions.GC()

	errs := router.ValidatePages()
	if len(errs) > 0 {
		for _, err := range errs {
			log.Print(err)
		}
		log.Fatal("Unable to continue due to template errors")
	}
	errs = router.ValidatePlugins()
	if len(errs) > 0 {
		for _, err := range errs {
			log.Print(err)
		}
		log.Fatal("Unable to continue due to plugin errors")
	}
	log.Fatal(router.Route(sessions))
}

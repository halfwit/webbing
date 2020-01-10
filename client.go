package main

import (
	"log"
	// Call our init functions to add our items
	_ "github.com/olmaxmedical/olmax_go/forms"
	_ "github.com/olmaxmedical/olmax_go/forms/doctor"
	_ "github.com/olmaxmedical/olmax_go/forms/patient"
	_ "github.com/olmaxmedical/olmax_go/pages"
	_ "github.com/olmaxmedical/olmax_go/pages/doctor"
	_ "github.com/olmaxmedical/olmax_go/pages/help"
	_ "github.com/olmaxmedical/olmax_go/pages/patient"
	_ "github.com/olmaxmedical/olmax_go/plugins"

	"github.com/olmaxmedical/olmax_go/router"
	"github.com/olmaxmedical/olmax_go/session"
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

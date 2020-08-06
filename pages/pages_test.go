package pages

import (
	"testing"
	
	"github.com/olmaxmedical/router"
	_ "github.com/olmaxmedical/pages/doctor"
	_ "github.com/olmaxmedical/pages/help"
	_ "github.com/olmaxmedical/pages/patient"
)

// TestPages largely makes sure we get back real data from every page
func TestPages(t *testing.T) {
	if e := router.RunPages(); e != nil {
		t.Error(e)
	}
}
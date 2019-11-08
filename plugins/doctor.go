package plugins

import (
	"github.com/olmaxmedical/olmax_go/router"
	"golang.org/x/text/message"
)

type doctor struct {
	Image     string
	AlmaMater string
	Name      string
	Residency string
	Current   string
	Country   string
	Specialty string
	Rate      string
}

func init() {
	b := &router.Plugin{
		Name:     "List doctors",
		Run:      ListDoctors,
		Validate: ValidateListDoctors,
	}
}

func ValidateListDoctors() error {
	return nil
}

func ListDoctors(p *message.Printer) []doctor {
	return []doctor{
		{
			Image:     "AbuzamzamMD.jpg",
			AlmaMater: "University of Southern California School of Medicine",
			Residency: "University of Southern California, San Diego. Internal Medicine Residency",
			Name:      "Mark Abuzamzam, MD",
			Current:   "Current Faculty at University of California Irvine Program Director",
			Country:   "United States of America",
			Specialty: "Internal Medicine and Addictions Medicine",
			Rate:      "0.0013 BTC",
		},
		{
			Image:     "WoodfinMD.jpg",
			Name:      "Martha Woodfin, MD",
			AlmaMater: "University Seoul School of Medicine",
			Residency: "University of Las Vegas Nevada, Pediatric Medicine Residency",
			Current:   "Current Staff at Mercy Hospital Jackson NC",
			Country:   "United States of America",
			Specialty: "Internal Medicine",
			Rate:      "0.0011 BTC",
		},
	}
}

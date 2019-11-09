package plugins

import (
	"github.com/olmaxmedical/olmax_go/router"
	"golang.org/x/text/message"
)

// ListDoctors - Bitmask to list doctors of in client country
const ListDoctors router.IncludeExtra = 2

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
		Run:      ListDocs,
		Validate: ValidateListDocs,
	}
	router.AddPlugin(b, ListDoctors)
}

// ValidateListDocs - Check that db entries exist
func ValidateListDocs() error {
	return nil
}

// ListDocs - Query db and return list of doctors in country
func ListDocs(p *message.Printer) map[string]interface{} {
	return map[string]interface{}{
		"Mark Abuzamzam, MD": &doctor{
			Image:     "AbuzamzamMD.jpg",
			AlmaMater: "University of Southern California School of Medicine",
			Residency: "University of Southern California, San Diego. Internal Medicine Residency",
			Name:      "Mark Abuzamzam, MD",
			Current:   "Current Faculty at University of California Irvine Program Director",
			Country:   "United States of America",
			Specialty: "Internal Medicine and Addictions Medicine",
			Rate:      "0.0013 BTC",
		},
		"Martha Woodfin, MD": &doctor{
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

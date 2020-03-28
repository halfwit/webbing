package plugins

import (
	"fmt"

	"github.com/olmaxmedical/olmax_go/router"
)

// Services - Bitmask to list services in native language
const Services router.PluginMask = 1 << 14

func init() {
	b := &router.Plugin{
		Name:     "specialties",
		Run:      ListServices,
		Validate: ValidateServices,
	}
	router.AddPlugin(b, Services)
}

// ValidateServices - Ensure the specialties entered exist in our map
func ValidateServices(r *router.Request) error {
	s := r.Request()
	var errs []string
	for _, entry := range s.PostFormValue("specialty") {
		switch string(entry) {
		case "acutepain":
		case "anasthesiology":
		case "bariatric":
		case "cardiology":
		case "chiropractic":
		case "chronic":
		case "critcare":
		case "dermatology":
		case "emergency":
		case "endocrinology":
		case "otolaringology":
		case "familymedicine":
		case "gastro":
		case "headneck":
		case "hematology":
		case "hepatology":
		case "hyperbaric":
		case "immunology":
		case "diseases":
		case "internal":
		case "neonatal":
		case "nephrology":
		case "neurology":
		case "neurosurgery":
		case "obstetrics":
		case "occupational":
		case "opthamology":
		case "orthopedics":
		case "palliative":
		case "pediatrics":
		case "podiatry":
		case "pulmonology":
		case "radiology":
		case "radiation":
		case "transplants":
			continue
		default:
			errs = append(errs, fmt.Sprintf("Unknown entry: %q\n", entry))
		}
		if len(errs) > 0 {
			return fmt.Errorf("%s", errs)
		}
	}
	return nil
}

// ListServices - return a list of native language representations of our medical fields
func ListServices(r *router.Request) map[string]interface{} {
	p := r.Printer()
	return map[string]interface{}{
		"label":          p.Sprint("Select specialty/specialties"),
		"acutepain":      p.Sprint("Acute Pain Medicine"),
		"anesthesiology": p.Sprint("Anesthesiology"),
		"bariatric":      p.Sprint("Bariatric Surgery"),
		"cardiology":     p.Sprint("Cardiology"),
		"chiropractic":   p.Sprint("Chiropractics"),
		"chronic":        p.Sprint("Chronic Pain"),
		"critcare":       p.Sprint("Critical Care"),
		"dermatology":    p.Sprint("Dermatology"),
		"emergency":      p.Sprint("Emergency Medicine"),
		"endocrinology":  p.Sprint("Endocrinology"),
		"otolaringology": p.Sprint("Ear Nose and Throat"),
		"familymedicine": p.Sprint("Family Medicine"),
		"gastro":         p.Sprint("Gastrointestinology"),
		"headneck":       p.Sprint("Head and Neck"),
		"hematology":     p.Sprint("Hematology and Oncology"),
		"hepatology":     p.Sprint("Hepatology"),
		"hyperbaric":     p.Sprint("Hyperbaric"),
		"immunology":     p.Sprint("Immunology"),
		"diseases":       p.Sprint("Infectious Diseases"),
		"internal":       p.Sprint("Internal Medicine"),
		"neonatal":       p.Sprint("Neonatology"),
		"nephrology":     p.Sprint("Nephrology"),
		"neurology":      p.Sprint("Neurology"),
		"neurosurgery":   p.Sprint("Neurosurgery"),
		"obstetrics":     p.Sprint("Obstetrics and Gynecology"),
		"occupational":   p.Sprint("Occupational Medicine"),
		"opthamology":    p.Sprint("Opthamology"),
		"orthopedics":    p.Sprint("Orthopedic Surgery"),
		"palliative":     p.Sprint("Palliative Care"),
		"pediatrics":     p.Sprint("Pediatrics"),
		"podiatry":       p.Sprint("Podiatry"),
		"pulmonology":    p.Sprint("Pulmonology"),
		"radiology":      p.Sprint("Radiology"),
		"radiation":      p.Sprint("Radiaton Oncology"),
		"transplants":    p.Sprint("Transplant Surgery"),
	}
}

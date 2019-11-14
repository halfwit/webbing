package plugins

import (
	"fmt"

	"github.com/olmaxmedical/olmax_go/router"
)

// Services - Bitmask to list services in native language
const Services router.PluginMask = 1 << 12

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
		case "anesthesiology":
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
		"acutepain":      p.Sprintf("Acute Pain Medicine"),
		"anesthesiology": p.Sprintf("Anesthesiology"),
		"bariatric":      p.Sprintf("Bariatric Surgery"),
		"cardiology":     p.Sprintf("Cardiology"),
		"chiropractic":   p.Sprintf("Chiropractics"),
		"chronic":        p.Sprintf("Chronic Pain"),
		"critcare":       p.Sprintf("Critical Care"),
		"dermatology":    p.Sprintf("Dermatology"),
		"emergency":      p.Sprintf("Emergency Medicine"),
		"endocrinology":  p.Sprintf("Endocrinology"),
		"otolaringology": p.Sprintf("Ear Nose and Throat"),
		"familymedicine": p.Sprintf("Family Medicine"),
		"gastro":         p.Sprintf("Gastrointestinology"),
		"headneck":       p.Sprintf("Head and Neck"),
		"hematology":     p.Sprintf("Hematology and Oncology"),
		"hepatology":     p.Sprintf("Hepatology"),
		"hyperbaric":     p.Sprintf("Hyperbaric"),
		"immunology":     p.Sprintf("Immunology"),
		"diseases":       p.Sprintf("Infectious Diseases"),
		"internal":       p.Sprintf("Internal Medicine"),
		"neonatal":       p.Sprintf("Neonatology"),
		"nephrology":     p.Sprintf("Nephrology"),
		"neurology":      p.Sprintf("Neurology"),
		"neurosurgery":   p.Sprintf("Neurosurgery"),
		"obstetrics":     p.Sprintf("Obstetrics and Gynecology"),
		"occupational":   p.Sprintf("Occupational Medicine"),
		"opthamology":    p.Sprintf("Opthamology"),
		"orthopedics":    p.Sprintf("Orthopedic Surgery"),
		"palliative":     p.Sprintf("Palliative Care"),
		"pediatrics":     p.Sprintf("Pediatrics"),
		"podiatry":       p.Sprintf("Podiatry"),
		"pulmonology":    p.Sprintf("Pulmonology"),
		"radiology":      p.Sprintf("Radiology"),
		"radiation":      p.Sprintf("Radiaton Oncology"),
		"transplants":    p.Sprintf("Transplant Surgery"),
	}
}

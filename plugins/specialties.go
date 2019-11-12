package plugins

import (
	"github.com/olmaxmedical/olmax_go/router"
)

// ListServices - Bitmask to list services in native language
const ListServices router.IncludeExtra = 2

func init() {
	b := &router.Plugin{
		Name:     "specialties",
		Run:      Specialties,
		Validate: ValidateSpecialties,
	}
	router.AddPlugin(b, ListServices)
}

// ValidateSpecialties - No-op
func ValidateSpecialties() error {
	return nil
}

// Specialties - return a list of native language representations of our medical fields
func Specialties(r *router.Request) map[string]interface{} {
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

package router

import (
	"golang.org/x/text/message"
)

// TODO: inverse function to get the actual specialty back from a whitelist
type Specialty struct {
	ID   string
	Name string
}

func specialties(p *message.Printer) []Specialty {
	return []Specialty{
		{"acutepain", p.Sprintf("Acute Pain Medicine")},
		{"anesthesiology", p.Sprintf("Anesthesiology")},
		{"bariatric", p.Sprintf("Bariatric Surgery")},
		{"cardiology", p.Sprintf("Cardiology")},
		{"chiropractic", p.Sprintf("Chiropractics")},
		{"chronic", p.Sprintf("Chronic Pain")},
		{"critcare", p.Sprintf("Chronic Pain")},
		{"dermatology", p.Sprintf("Dermatology")},
		{"emergency", p.Sprintf("Emergency Medicine")},
		{"endocrinology", p.Sprintf("Endocrinology")},
		{"otolaringology", p.Sprintf("Ear Nose and Throat")},
		{"familymedicine", p.Sprintf("Family Medicine")},
		{"gastro", p.Sprintf("Gastrointestinology")},
		{"headneck", p.Sprintf("Head and Neck")},
		{"hematology", p.Sprintf("Hematology and Oncology")},
		{"hepatology", p.Sprintf("Hepatology")},
		{"hyperbaric", p.Sprintf("Hyperbaric")},
		{"immunology", p.Sprintf("Immunology")},
		{"diseases", p.Sprintf("Infectious Diseases")},
		{"internal", p.Sprintf("Internal Medicine")},
		{"neonatal", p.Sprintf("Neonatology")},
		{"nephrology", p.Sprintf("Nephrology")},
		{"neurology", p.Sprintf("Neurology")},
		{"neurosurgery", p.Sprintf("Neurosurgery")},
		{"obstetrics", p.Sprintf("Obstetrics and Gynecology")},
		{"occupational", p.Sprintf("Occupational Medicine")},
		{"opthamology", p.Sprintf("Opthamology")},
		{"orthopedics", p.Sprintf("Orthopedic Surgery")},
		{"palliative", p.Sprintf("Palliative Care")},
		{"pediatrics", p.Sprintf("Pediatrics")},
		{"podiatry", p.Sprintf("Podiatry")},
		{"pulmonology", p.Sprintf("Pulmonology")},
		{"radiology", p.Sprintf("Radiology")},
		{"radiation", p.Sprintf("Radiaton Oncology")},
		{"transplants", p.Sprintf("Transplant Surgery")},
	}
}

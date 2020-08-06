package patient

import (
	"github.com/olmaxmedical/plugins"
	"github.com/olmaxmedical/router"
	"golang.org/x/text/message"
)

func init() {
	b := &router.Page{
		Access: router.PatientAuth,
		CSS:    "",
		Path:   "patient/symptoms",
		Data:   symptoms,
		Extra:  plugins.FormErrors,
	}
	router.AddPage(b)
}

func symptoms(p *message.Printer) map[string]interface{} {
	return map[string]interface{}{
		"title":                p.Sprintf("Olmax Medical | Create Profile"),
		"createHeader":         p.Sprintf("Create Patient Profile"),
		"profileHeader":        p.Sprintf("Create Patient Profile"),
		"formHeader":           p.Sprintf("Please submit some information regarding your consult."),
		"birthdate":            p.Sprintf("When were you born?"),
		"asab":                 p.Sprintf("Which sex were you assigned at birth?"),
		"visitReason":          p.Sprintf("Please give a brief statement regarding the main reason you would like to see your doctor:"),
		"symptomStart":         p.Sprintf("When did your symptoms start?"),
		"symptomArea":          p.Sprintf("Where are your symptoms located? <i>part of your body</i>"),
		"symptomDuration":      p.Sprintf("How long have these symptoms lasted?"),
		"symptomDescription":   p.Sprintf("How would you characterize your symptoms? <i>Sharp, Dull, Ache, twisting, ets.</i>"),
		"symptomAugment":       p.Sprintf("What makes your symptoms better, and What makes them worse?"),
		"symptomProliferation": p.Sprintf("Does your pain travel or radiate to another part of your body?"),
		"symptomTreatment":     p.Sprintf("Have you taken any medications for these symptoms and how much have they worked?"),
		"feversChills":         p.Sprintf("Any fevers or Chills?"),
		"wtGainLoss":           p.Sprintf("Any weight gain or weight loss?"),
		"vision":               p.Sprintf("Any changes in vision?"),
		"lung":                 p.Sprintf("Any lung issues?"),
		"heart":                p.Sprintf("Any heart problems?"),
		"bowel":                p.Sprintf("Any intestinal problems?"),
		"renal":                p.Sprintf("Any kidney problems?"),
		"musSkel":              p.Sprintf("Any problems with muscles or bones?"),
		"neuro":                p.Sprintf("Any nervous system problmes? <i>Strokes</i>"),
		"psych":                p.Sprintf("Any psychiatric problems? <i>Depression, anxiety</i>"),
		"male":                 p.Sprintf("Male"),
		"female":               p.Sprintf("Female"),
		"legend":               p.Sprintf("Please check any of the following if you have experienced in the last 6 weeks:"),
	}
}

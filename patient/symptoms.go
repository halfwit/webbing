package forms

import (
	"net/http"
	"time"

	"github.com/albrow/forms"
	"github.com/olmaxmedical/plugins"
	"github.com/olmaxmedical/router"
	"golang.org/x/text/message"
)

func init() {
	b := &router.Form{
		Access:    router.PatientAuth,
		Path:      "patient/symptoms",
		Validator: symptoms,
		After:     plugins.EmailForm,
		Redirect:  "patient/profile.html",
	}
	router.AddPost(b)
}

func symptoms(r *http.Request, p *message.Printer) []string {
	var errors []string
	data, err := forms.Parse(r)
	if err != nil {
		errors = append(errors, p.Sprint("Internal server error"))
		return errors
	}
	val := data.Validator()
	// TODO(halfwit): Date must be in a reasonable range
	val.Require("bday").Message(p.Sprint("Birth date required"))
	_, err = time.Parse("2006-01-02T15:04:05", r.Form.Get("bday"))
	if err != nil {
		val.AddError("bday", p.Sprint("Invalid birth date"))
	}
	val.Require("onset").Message(p.Sprint("Please enter the date and time your symptoms started"))
	_, err = time.Parse("2006-01-02T15:04:05", r.Form.Get("onset"))
	if err != nil {
		val.AddError("bday", p.Sprint("Invalid date"))
	}
	val.Require("gender").Message(p.Sprint("Please select a biological gender"))
	if r.PostFormValue("gender") != "male" && r.PostFormValue("gender") != "female" {
		val.AddError("gender", p.Sprint("Invalid selection for gender"))
	}
	val.GreaterOrEqual("duration", 0).Message(p.Sprint("Invalid value entered for how long symptoms have lasted"))
	val.Require("reason").Message(p.Sprint("Please provide the reason for visit"))
	val.Require("location").Message(p.Sprint("Please list the area the symptom(s) appear"))
	val.Require("characteristic").Message(p.Sprint("Please provide a description of your symptoms"))
	val.Require("aggreAlevi").Message(p.Sprint("Please note anything which improves/worsens your symptoms"))
	for _, i := range []string{
		"feversChills",
		"wtGainLoss",
		"vision",
		"lung",
		"heart",
		"bowel",
		"renal",
		"musSkel",
		"neuro",
		"psych",
	} {
		sel, ok := r.Form[i]
		if !ok {
			val.AddError(i, p.Sprintf("No selection for %s", i))
			continue
		}
		if sel[0] == "Yes" || sel[0] == "yes" || sel[0] == "no" || sel[0] == "No" {
			continue
		}
		val.AddError(i, p.Sprintf("Invalid selection for %s", i))
	}
	r.Form["pagetitle"] = []string{"Client symptoms"}
	if val.HasErrors() {
		errors = append(errors, val.Messages()...)
	}
	return errors
}

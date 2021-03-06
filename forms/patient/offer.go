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
		Path:      "patient/offer",
		Validator: offer,
		After:     plugins.Search | plugins.Services, //|plugins.Offer
		Redirect:  "results.html",
	}
	router.AddPost(b)
}

func offer(r *http.Request, p *message.Printer) []string {
	var errors []string
	data, err := forms.ParseMax(r, r.ContentLength)
	if err != nil {
		errors = append(errors, p.Sprint("Internal server error"))
		return errors
	}
	val := data.Validator()
	val.Require("Amount").Message(p.Sprint("Please enter a target rate (Bitcoin/15min)"))
	bcu := data.GetFloat("Amount")
	if 0.0 > bcu || bcu > 1.0 {
		val.AddError("Amount", p.Sprint("BTC/15min rate out of range"))
	}
	val.Require("startDate").Message(p.Sprint("Start date required"))
	_, err = time.Parse("2006-01-02T15:04:05", r.Form.Get("startDate"))
	if err != nil {
		val.AddError("startDate", p.Sprint("Invalid start-date entered"))
	}

	val.Require("endDate").Message(p.Sprint("End date required"))
	_, err = time.Parse("2006-01-02T15:04:05", r.Form.Get("endDate"))
	if err != nil {
		val.AddError("endDate", p.Sprint("Invalid end-date entered"))
	}

	if val.HasErrors() {
		errors = append(errors, val.Messages()...)
	}
	return errors
}

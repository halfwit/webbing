package forms

import (
	"fmt"
	"net/http"

	"github.com/albrow/forms"
	"github.com/olmaxmedical/olmax_go/plugins"
	"github.com/olmaxmedical/olmax_go/router"
	"golang.org/x/text/message"
)

func init() {
	b := &router.Form{
		Access:    router.GuestAuth,
		Path:      "doctor/application",
		Validator: Application,
		Redirect:  "/index.html",
		After:     plugins.EmailForm | plugins.Countries | plugins.Services | plugins.FormToken,
	}
	router.AddPost(b)
}

// Application - olmaxmedical.com/doctor/application.html
func Application(r *http.Request, p *message.Printer) []string {
	var errors []string
	data, err := forms.Parse(r)
	if err != nil {
		errors = append(errors, "Internal server error")
		return errors
	}
	val := data.Validator()
	val.Require("gender").Message(p.Sprint("Please select a biological gender"))
	if r.PostFormValue("gender") != "male" && r.PostFormValue("gender") != "female" {
		val.AddError("gender", p.Sprint("Invalid selection for gender"))
	}
	val.RequireFile("cv").Message(p.Sprint("Empty or missing CV"))
	val.AcceptFileExts("cv", "application/msword,applicationvnd.openxmlformats-officedocument.wordprocessingml.document,application/pdf").Message(p.Sprint("Unsupported filetype for CV"))
	val.RequireFile("diploma").Message(p.Sprint("Empty or missing Diploma/Board Certification"))
	val.AcceptFileExts("cv", "application/msword,applicationvnd.openxmlformats-officedocument.wordprocessingml.document,application/pdf").Message(p.Sprint("Unsupported filetype for Diploma/Board Certification"))
	for i := 1; i < 12; i++ {
		num := fmt.Sprintf("q%d", i)
		sel, ok := r.Form[num]
		if !ok {
			val.AddError(num, p.Sprintf("No selection for question %d", i))
			continue
		}
		if sel[0] == "Yes" || sel[0] == "yes" || sel[0] == "no" || sel[0] == "No" {
			continue
		}
		val.AddError(num, p.Sprintf("Invalid selection for question %d", i))
	}
	val.Require("email").Message(p.Sprintf("Valid email required"))
	val.MatchEmail("email").Message(p.Sprintf("Invalid email"))
	val.Require("name").Message(p.Sprintf("Full name required"))
	val.MinLength("name", 2).Message(p.Sprintf("Full name must be at least 2 characters"))
	if r.PostFormValue("redFlag") != "on" {
		val.AddError("redFlag", p.Sprint("Invalid selection for confirm element"))
	}
	if val.HasErrors() {
		errors = append(errors, val.Messages()...)
	}
	r.Form["pagetitle"] = []string{"Application for doctor"}
	r.Form["sendto"] = []string{"olmaxmedical@gmail.com"}
	delete(r.Form, "redFlag")
	return errors
}

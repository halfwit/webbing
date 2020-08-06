package forms

import (
	"bytes"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"testing"

	"github.com/olmaxmedical/forms/util"
	"golang.org/x/text/message"
)

func TestApplication(t *testing.T) {
	fields := map[string]string{
		"qs":      "no",
		"gender":  "male",
		"email":   "foo@bar.ca",
		"name":    "Doctor Octopus",
		"redFlag": "on",
	}

	if e := testApplication(t, fields); e != nil {
		t.Error(e)
	}

	fields["gender"] = "pineapple"
	if e := testApplication(t, fields); e == nil {
		t.Error("invalid field accepted")
	}

	fields["gender"] = "male"
	fields["email"] = "foo@bar@ca"
	if e := testApplication(t, fields); e == nil {
		t.Error("invalid field accepted")
	}

	fields["email"] = "foo@bar.ca"
	fields["qs"] = "true"
	if e := testApplication(t, fields); e == nil {
		t.Error("invalid field accepted")
	}
}

func testApplication(t *testing.T, fields map[string]string) error {
	var reqBody bytes.Buffer

	mpw := multipart.NewWriter(&reqBody)
	files := map[string]string{
		"cv":      "resume.pdf",
		"diploma": "certificate.pdf",
	}

	for key, value := range files {
		if e := util.WriteFile(mpw, key, "../resources/"+value); e != nil {
			panic(e)
		}
	}

	for key, value := range fields {
		if key == "qs" {
			for i := 0; i < 12; i++ {
				key = fmt.Sprintf("q%d", i)
				if e := mpw.WriteField(key, value); e != nil {
					panic(e)
				}
			}
			continue
		}

		if e := mpw.WriteField(key, value); e != nil {
			panic(e)
		}
	}

	request := util.BuildMultiRequest(mpw, &reqBody)
	printer := message.NewPrinter(message.MatchLanguage("en"))
	return runTest(request, printer)
}

func runTest(request *http.Request, printer *message.Printer) error {
	for _, err := range application(request, printer) {
		switch err {
		case "unsupported filetype for cv":
		case "unsupported filetype for diploma":
		default:
			return errors.New(err)
		}

	}

	return nil
}

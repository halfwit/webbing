package forms

import (
	"bufio"
	"net/http"
	"strings"
	"testing"
)

// This isn't quite right yet, we'll get a working test soon.
var req = `POST /doctor/application.html HTTP/1.1
Host: localhost
User-Agent: Mozilla/5.0
Accept: */*
Connection: keep-alive
Cache-Control: no-cache
Content-Length: 2028
Content-Type: multipart/form-data; boundary=--------------------

--------------------------4ee5f9ad0b2d7899
Content-Disposition: form-data; name="Specialty"

bariatric
--------------------------4ee5f9ad0b2d7899
Content-Disposition: form-data; name="country"

Albania
--------------------------4ee5f9ad0b2d7899
Content-Disposition: form-data; name="email"

mee%40foo.ca
--------------------------4ee5f9ad0b2d7899
Content-Disposition: form-data; name="gender"

male
--------------------------4ee5f9ad0b2d7899
Content-Disposition: form-data; name="name"

mememe
--------------------------4ee5f9ad0b2d7899
Content-Disposition: form-data; name="pagetitle"

Application+for+doctor
--------------------------4ee5f9ad0b2d7899
Content-Disposition: form-data; name="q1"

No
--------------------------4ee5f9ad0b2d7899
Content-Disposition: form-data; name="q10"

No
--------------------------4ee5f9ad0b2d7899
Content-Disposition: form-data; name="q11"

No
--------------------------4ee5f9ad0b2d7899
Content-Disposition: form-data; name="q2"

No
--------------------------4ee5f9ad0b2d7899
Content-Disposition: form-data; name="q3"

No
--------------------------4ee5f9ad0b2d7899
Content-Disposition: form-data; name="q4"

No
--------------------------4ee5f9ad0b2d7899
Content-Disposition: form-data; name="q5"

No
--------------------------4ee5f9ad0b2d7899
Content-Disposition: form-data; name="q6"

No
--------------------------4ee5f9ad0b2d7899
Content-Disposition: form-data; name="q7"

No
--------------------------4ee5f9ad0b2d7899
Content-Disposition: form-data; name="q8"

No
--------------------------4ee5f9ad0b2d7899
Content-Disposition: form-data; name="q9"

No
--------------------------4ee5f9ad0b2d7899
Content-Disposition: form-data; name="sendto"

olmaxmedical%40gmail.com
--------------------------4ee5f9ad0b2d7899
Content-Disposition: form-data; name="cv"

resume.pdf
--------------------------4ee5f9ad0b2d7899
Content-Disposition: form-data; name="form"

resume.pdf
--------------------------4ee5f9ad0b2d7899`

func TestApplication(t *testing.T) {
	// Build a POST request with a map of entries
	rd := bufio.NewReader(strings.NewReader(req))

	request, err := http.ReadRequest(rd)
	if err != nil {
		t.Errorf("test design error: bad request: %v", err)
		return
	}

	if e := request.ParseMultipartForm(request.ContentLength); e != nil {
		t.Error(e)
	}
	/*
		printer := message.NewPrinter(message.MatchLanguage("en"))

		resp := application(request, printer)
		if len(resp) > 0 {
			for _, err := range resp {
				t.Errorf("%s", err)
			}
		}
	*/
}

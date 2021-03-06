package email

import (
	"bytes"
	"log"
	"mime/multipart"
	"net/mail"
	"net/smtp"
	"text/template"

	"github.com/scorredoira/email"
	"golang.org/x/text/message"
)

const tmpl = `
<!DOCTYPE html>
<html>
	<title>{{index .pagetitle 0}}</title>
<head>
	
</head>
<body>
	<h2>{{index .pagetitle 0}}</h2>
	<ul>
	{{range $key, $value := . }}
		{{if eq $key "pagetitle"}}
		{{else if eq $key "country" }}
			<li>countries: {{range $value}}{{.}} {{end}}</li>
		{{else if eq $key "specialty" }}
			<li>specialties: {{range $value}}{{.}} {{end}}</li>
		{{else}}
			{{$val := len $value}}
			{{if eq $val 2}}
				<li>{{$key}}: {{index $value 1 }}</li>
			{{else}}
				<li>{{$key}}: {{index $value 0 }}</li>
			{{end}}
		{{end}}
	{{end}}
	</ul>
</body>
</html>
`

var t *template.Template

func init() {
	t = template.Must(template.New("email").Parse(tmpl))
}

// SendForm - Fill in the template and send it out from our email address
func SendForm(form map[string][]string, p *message.Printer, attachments map[string]multipart.File) {
	var body bytes.Buffer

	// Stash the address we want to send to
	address := form["sendto"][0]
	delete(form, "sendto")

	if e := t.Execute(&body, form); e != nil {
		log.Println(e)
		return
	}

	m := email.NewHTMLMessage("Form contents", body.String())
	m.From = mail.Address{
		Name:    "From",
		Address: gmail,
	}

	m.AddTo(mail.Address{
		Name:    "To",
		Address: address,
	})

	for name, buff := range attachments {
		var attc bytes.Buffer

		attc.ReadFrom(buff)

		if err := m.AttachBuffer(name, attc.Bytes(), false); err != nil {
			log.Println(err)
		}
	}

	auth := smtp.PlainAuth("", gmail, pw, addr)
	if err := email.Send(addr+":587", auth, m); err != nil {
		log.Println(err)
	}
}

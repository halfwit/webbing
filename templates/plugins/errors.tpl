{{define "errors"}}
{{range .}}
<p style="color: red" class="errtext">{{.}}</p>
{{end}}
{{end}}
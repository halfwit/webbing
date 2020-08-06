{{define "country"}}
<label for="country">{{.label}}</label>
<select name="country" id="country" multiple required>
    {{range $id, $name := .}}
	<option value="{{$id}}">{{$name}}</option>
    {{end}}
</select>
{{end}}

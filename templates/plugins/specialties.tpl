{{define "specialty"}}				
<label for="Specialty">{{.label}}</label>
<select name="Specialty" id="specialty" multiple required>
    {{range $id, $name := .}}
	<option value="{{$id}}">{{$name}}</option>
	{{end}}		
</select>
{{end}}
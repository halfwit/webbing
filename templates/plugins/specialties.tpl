{{define "specialty"}}				
<label for="Specialty">{{.specialty}}</label>
	<select name="Specialty">
        {{range $id, $name := .}}
		<option value="{{$id}}">{{$name}}</option>
	    {{end}}		
    </select>
{{end}}
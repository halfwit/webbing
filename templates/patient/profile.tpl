{{define "content"}}
	<main>
 	<h2>{{.greetingHeader}}{{.username}}</h2>	    
	<fieldset id="find-dr">
	<form method=POST action="/getappointments">
	    <legend>{{.legend}}</legend>
	    <p>{{.specialty}}</p>
	    <select name="specialty">
	    {{range .specialties}}
	        <option value="{{.ID}}">{{.Name}}</option>
	    {{end}}
	    </select>
	    <p>{{.country}}</p>
	    <select name="country">
	    {{range .countrylist}}
		<option value="{{.ID}}">{{.Name}}</option>
            {{end}}
	    </select>
	    <legend>{{.apptLegend}}</legend>
		<label for="startDate">{{.from}}</label>
		<input id="startDate" type="date" name="startDate" required>
		<label for="endDate">{{.to}}</label>
		<input id="endDate" type="date" name="endDate" required>		
	    <button>{{.search}}</button>
	</form>
	</fieldset>
	</main>
{{end}}
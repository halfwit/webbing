{{define "content"}}
	<main>
	 <h2>{{.mainHeader}}</h2>
	 <form action="findspecialty">
		{{template "specialty" .specialties}}
		<br/>
		<label for="Amount">{{.bcu}} <i>({{.rate}})</i></label>
		<input type="text" name="Amount" size="15" maxlength="10" />
	    	<h3>{{.dates}}</h3>
		<label for="startDate">{{.from}}</label>
		<input id="startDate" type="date" name="startDate" required>
		<label for="endDate">{{.to}}</label>
		<input id="endDate" type="date" name="endDate" required>
		<br/>
		<button>{{.deploy}}</button>
	</form>
	</main>
{{end}}
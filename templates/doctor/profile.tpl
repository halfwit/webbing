{{define "content"}}
	<main>
 	<h2>{{.greetingHeader}}{{.username}}</h2>
	<form method=POST>
		<legend for="dates">{{.appLegend}}</legend>
		<fieldset id="dates">
			<!-- TODO: review better ways to set up chunks of availability
			     set up a recurring chunk for a range of hours with weekday checkboxes
			     then create a method for editing single days
			     this should be flexible enough to get out of the way generally
			     but still facilitate granular control when necessary
			-->
			<label for="startDate">{{.from}}</label>
			<input id="startDate" name="startDate" type="datetime-local" required>
			<br>
			<label for="endDate">{{.to}}</label>
			<input id="startDate" name="startDate" type="datetime-local" required>
			<p>{{.bcu}}</p>
			<label for="BTCperU">BTC:</label>
			<input id="BTCperU" id="BTCperU" type="number" min="0.0012" max="1.0" step="0.00001" required>
		</fieldset>
		<button>{{.create}}</button>
		<input type="hidden" name="token" value="{{.token}}">
		<input type="reset">
	</form>
	</main>
{{end}}
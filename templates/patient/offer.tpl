{{define "content"}}
	<main>
	 <h2>{{.mainHeader}}</h2>
 	<table>
		<th scope="row">{{.specialtyHeader}}</th>
		<th></th>
		<td>
		    <form action="findspecialty">
			<select name="Specialty">
			{{range .specialties}}
			    <option value="{{.ID}}">{{.Name}}</option>
			{{end}}
			</select>
		    </form>
		</td>
	    <!-- needs backend go files -->
	    <form action=" .go file">
	    <tr>
		<th scope="row">{{.bcu}}</th>
		<th></th>
		<td><input type="text" name="Amount" size="15" maxlength="10" /> </td>
	    </tr>
		<th scope="row">{{.dates}}</th>
		<th></th>
		<td>
		    <!-- needs go files -->
		    <form action="some go file" method="get"> 
			<label for="startDate">{{.from}}</label>
			<input id="startDate" type="date" name="startDate" required>
			<label for="endDate">{{.to}}</label>
			<input id="endDate" type="date" name="endDate" required>
		    </form>
		</td>
		<tr>
			<th></th>
			<th></th>
			<th><button>{{.deploy}}</button></th>
		</tr>
	</table>
	<hr/>
	</main>
{{end}}
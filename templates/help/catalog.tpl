{{define "content"}}
	<main>
	<h1>{{.mainHeader}}<h2>
<table border="1" cellpadding="8" cellspacing="2">
<tr bgcolor="gray">
	<th colspan="2">Profile</th>
	<th>Country</th><th>Specialty</th><th>Rate <i>BTC</i></th></th>
	{{range .doctors}}
  	<tr valign="top">
		<td><img src="../images/{{.Image}}" /><a href="invoice.html"><h4>{{.Name}}</h4></a></td>
	<td><p>{{.AlmaMater}}</p>
		<p>{{.Residency}}</p>
		<p>{{.Current}}</p>
	<td>{{.Country}}</td>
	<td>{{.Specialty}}</td>
	<td>{{.Rate}}</td>
	{{end}}
	</table>
	</main>
{{end}}
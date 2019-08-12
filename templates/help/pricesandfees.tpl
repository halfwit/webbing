{{define "content"}}
	<main>
	<h1>{{.mainHeader}}</h1>
	<h2>{{.priceHeader}}</h2>
		<p>{{.priceBody}}</p>
	<br>
	<h4>{{.determineHeader}}</h4>
		<p>{{.determineBody}}<p>
	<br>
	<h4>{{.chargedHeader}}</h4>
		<p>{{.chargedBody}}</p>
	<br>
	<h4>{{.currencyHeader}}</h4>
		<p>{{.currencyBody}}</p>
	<br>
	<h4>{{.blockHeader}}</h4> 
	</main>
{{end}}
{{define "content"}}
	<main>
	<h2>{{.mainHeader}}</h2>
	<div align="center">
		<h1>{{.funds}}</h1>
			<p>{{.current}}</p>
		</hr>
	</div>
	<div align="center">
		<h2>{{.deposit}}</h2>
		<img src="../images/bitcoinRQ.jpg" alt="[redacted]" width="155" height="155"/>	
			<h3>[redacted]</h3>
		<h4>{{.onlyHeader}}</h4>
		<p>{{.onlyBody}}</p>
	</main>
{{end}}

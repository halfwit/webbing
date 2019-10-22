{{define "content"}}
	<main>
	<h1>{{.mainHeader}}</h1>
	<h2>{{.contactHeader}}</h2>
		<p>{{.contactBody}}</p>
	<br>
	<h4>{{.scheduleHeader}}</h4>
		<p>{{.scheduleBody}}<p>
	<br>
	<h4>{{.chargedHeader}}</h4>
		<p>{{.chargedBody}}</p>
	<br>
	<h4>{{.anycurrHeader}}</h4>
		<p>{{.anycurrBody}}</p>
	<br>
	<h4>{{.blocksHeader}}</h4> 
                <p>{{.blocksBody}} <a href="https://www.torproject.org">Tor</a></p>
	</main>
{{end}}
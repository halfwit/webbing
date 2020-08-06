{{define "content"}}
	<main>
	<h1>{{.mainHeader}}</h1>
	<h2>{{.verifyHeader}}</h2>
		<p>{{.verifyBody}}</p>
	<h4>{{.phoneHeader}}</h4>
		<p>{{.phoneBody}}</p>
	<h4>{{.noNoteHeader}}</h4>
		<p>{{.noNoteBody}}<a href="{{.basedir}}/appointmentRequests.html">Appointment Requests</a></p>

	<h4>{{.blockHeader}}</h4> 
                <p>{{.blockBody}}<a href="https://www.torproject.org">Tor</a></p>
	</main>
{{end}}
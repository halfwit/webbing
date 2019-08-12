{{define "content"}}
	<main>
	<h1>{{.mainHeader}}</h1>
	<h2>{{.verifyHeader}}</h2>
		<p>{{.verifyBody}}</p>
	<h4>{{.phoneHeader}}</h4>
		<p>{{.phoneBody}}</p>
	<h4>{{.noNoteHeader}}</h4>
		<p>{{.noNoteBody}}</p>

	<h4>{{.blockHeader}}</h4> 
                <p>{{.blockBody}}</p>
	</main>
{{end}}
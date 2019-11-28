{{define "content"}}
	<main>
	    <h2>{{.mainHeader}}</h2>
	    <h3>{{.earn}}</h3>
	    <form method="post">
	        <p>{{.specialtyHeader}}</p>
	        <select name="specialty">
		    {{range $id, $name := .specialties}}
		    <option value="{{$id}}">{{$name}}</option>
		    {{end}}
		</select>
	        <button>{{.getStartedHeader}}</button>
	    </form>
		<!--TODO change link from TODO.html-->
	    <a href="TODO.html">{{.getStartedLink}}</a> 
	    <h3>{{.providerWhy}} </h3>
	    <p>{{.whyText}}</p>
	    <h3>{{.control}}</h3>
	    <p>{{.controlText}}</p>
	    <h3>{{.everyStep}}</h3>
	    <p>{{.everyStepText}}</p>
	    <h1><strong>{{.provider}}</strong></h1>
	    <h1><strong>1</strong></h1>
	    <h4>{{.createProfile}}</h4>
	    <p>{{.createProfileText}}</p>
	    <h1><strong>2</strong></h1>
	    <h4>{{.welcomePatient}}</h4>
	    <p>{{.welcomePatientText}}</p>
	    <h1><strong>3</strong></h1>
	    <h4>{{.getPaid}}</h4>
	    <p>{{.getPaidText}}</p> 
	    <h1><strong>{{.safety}}</strong></h1>
	    <h3>{{.trust}}</h3>
	    <p>{{.trustText}}</p>
	</main>
{{end}}
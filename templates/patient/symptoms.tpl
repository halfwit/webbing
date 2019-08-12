{{define "content"}}
	<main>
	<h2>{{.createHeader}}</h2>
	<hr/>
	<fieldset id="ptMedHx"> <!-- Hx=history, pt=patient, Med,Surg,Fam, Soc are obvious -->
	<h2>{{.formHeader}}</h2>
	<form action="post">
		<p>{{.birthdate}}</p>
		<input type="date" name="bday">
		<br>
		<p>{{.gender}}</p>
		<br>
		<input type="radio" name="gender" value="male">{{.male}}<br>
		<input type ="radio" name="gender" value="female">{{.female}}<br>
	<br>
	<p>{{.visitReason}}</p>
		<textarea row="6" cols="50">
		</textarea>
		<br>    
	<p>{{.symptomStart}}</p>
		<input type="date" name="onset">
		<br>
	<p>{{.symptomArea}}</p>
		<input type="text" id="location">
		<br>
	<p>{{.symptomDuration}}</p>
		<input type="number" id="duration">
		<br>
	<p>{{.symptomDescription}}</p> 
		<input type="text" id="characteristic" >
		<br>
	<p>{{.symptomAugment}}</p>
		<input type="text" id="aggreAlevi" >
		<br>
	<p>{{.symptomProliferation}}</p>
		<input type="text" id="radiate">
		<br>
	<p>{{.symptomTreatment}}</p>
		<textarea row="6" cols="50">
		</textarea>
	<fieldset>
	<legend>{{.legend}}</legend>
	<!--"ros" means Review of Systems -->
		<input type="radio" name="ros" id="feversChills" value="feversChills" />
		<label for="feversChills">{{.feversChills}}</label><br />
		<input type="radio" name="ros" id="wtGainLoss" value="wtGainLoss" />
		<label for="wtGainLoss">{{.wtGainLoss}}</label><br />
		<input type="radio" name="ros" id="vison" value="vision" />
		<label for="vision">{{.vision}}</label><br />
		<input type="radio" name="ros" id="lung" value="lung" />
		<label for="lung">{{.lung}}</label><br />
		<input type="radio" name="ros" id="heart" value="heart" />
		<label for="heart">{{.heart}}</label><br />
		<input type="radio" name="ros" id="bowel" value="bowel" />
		<label for="bowel">{{.bowel}}</label><br />
		<input type="radio" name="ros" id="renal" value="renal" />
		<label for="renal">{{.renal}}</label><br />
		<input type="radio" name="ros" id="musSkel" value="musSkel" />
		<label for="musSkel">{{.musSkel}}</label><br />
		<input type="radio" name="ros" id="neuro" value="neuro" />
		<label for="neuro">{{.neuro}}</label><br />
		<input type="radio" name="ros" id="psych" value="psych">
		<label for="psych">{{.psych}}<br />
	</fieldset>
<!-- mothra file:///usr/glenda/Olmax/ptProfile.html -->
		<br>
		<input type="submit" value="submit">
	</form>
	</main>
{{end}}
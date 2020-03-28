{{define "content"}}
	<!-- This will be emailed directly to the physician on completion
	     it is locally validated, but never locally stored
	 -->
	<main>
	<h2>{{.createHeader}}</h2>
	<hr/>
	<h2>{{.formHeader}}</h2>
	<form action="post">
		<label for="bday">{{.birthdate}}</label>
		<input type="date" name="bday"/>
		<label for="gender">{{.gender}}</label>
		<input type="radio" name="gender" value="male"/>{{.male}}
		<input type ="radio" name="gender" value="female"/>{{.female}}
		<label for="reason">{{.visitReason}}</label>
		<textarea name="reason" row="6" cols="50"></textarea>
		<label for="onset">{{.symptomStart}}</label>
		<input type="date" name="onset"/>
		<label for="location">{{.symptomArea}}</label>
		<input type="text" id="location"/>
		<laber for="duration">{{.symptomDuration}}</label>
		<input type="number" min="0" id="duration"/>
		<label for="characteristic">{{.symptomDescription}}</label> 
		<input type="text" id="characteristic"/>
		<label for="aggreAlevi">{{.symptomAugment}}</label>
		<input type="text" id="aggreAlevi"/>
		<label for="radiate">{{.symptomProliferation}}</label>
		<input type="text" id="radiate"/>
		<label for="treatments">{{.symptomTreatment}}</label>
		<textarea name="treatments" row="6" cols="50"></textarea>
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
			<label for="psych">{{.psych}}<br/>
		</fieldset>
		<input type="submit" value="submit">
	</form>
	</main>
{{end}}
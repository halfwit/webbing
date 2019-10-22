{{define "content"}}
	<main>
	        {{template "errors" .errors}}
		<label for="application">{{.offer}}</label>
		<form id="application" method="post" enctype="multipart/form-data" boundary="VAL">
			<h4>{{.area}}</h4>
			<fieldset>
				<label for="country">{{.doccountry}}</label>
				<select name="country" id="country" multiple required>
		    		{{range .countrylist}}
					<option value="{{.ID}}">{{.Name}}</option>
  	          		{{end}}
				</select>
				<br><!-- clean up the formatting on some browsers -->
				<label for="specialty">{{.docspecial}}</label>	   
		 		<select name="specialty" id="specialty" multiple required>
				{{range .specialties}}
					<option value="{{.ID}}">{{.Name}}</option>
				{{end}}
				</select>
			</fieldset>
			<h4>{{.gender}}</h4>
			<fieldset>
				<!-- for validation, we use the `name` to get the value associated with the radio cluster. ID is to associate labels -->
				<label for="male">{{.male}}</label>
				<input type="radio" name="gender" id="male" value="male" required>
				<label for="female">{{.female}}</label>
				<input type ="radio" name="gender" id="female" value="female" required>
			</fieldset>
			<h4>{{.documents}}</h4>
			<fieldset>
				<label for="cv">{{.cv}}</label>
				<!-- TODO: Reasonable format list -->
				<input type="file" id="cv" name="cv" accept=".pdf,.doc,.docx,.epub,.rtf,application/msword,applicationvnd.openxmlformats-officedocument.wordprocessingml.document" required>
				<br>
				<label for="diploma">{{.diploma}}</label>
				<input type="file" id="diploma" name="diploma" accept="application/msword,applicationvnd.openxmlformats-officedocument.wordprocessingml.document,application/pdf,.pdf,.doc,.docx,.epub,.rtf," required multiple>
			</fieldset>
			<!-- To use `required`, both input elements must have the same name -->
			<!-- for validation, we will use the `name` to get the value associated with the radio cluster -->
			<h4>I</h4>
			<fieldset id="q1">
				<p>{{.q1}}</p>
				<label for="q1yes">{{.yes}}</label>
				<input type="radio" name="q1" id="q1yes" value="Yes" required>
				<label for="q1no">{{.no}}</label>
				<input type ="radio" name="q1" id="q1no" value="No" required>
			</fieldset>
			<h4>II</h4>
			<fieldset id="q2">
				<p>{{.q2}}</p>
				<label for="q2yes">{{.yes}}</label>
				<input type="radio" name="q2" id="q2yes" value="Yes" required>
				<label for="q2no">{{.no}}</label>
				<input type="radio" name="q2" id="q2no" value="No" required>
			</fieldset>
			<h4>III</h4>
			<fieldset id="q3">
				<p>{{.q3}}</p>
				<label for="q3yes">{{.yes}}</label>
				<input type="radio" name="q3" id="q3yes" value="Yes" required>
				<label for="q3no">{{.no}}</label>
				<input type ="radio" name="q3" id="q3no" value="No" required>
			</fieldset>
			<h4>IV</h4>
			<fieldset id="q4">
				<p>{{.q4}}</p>
				<label for="q4yes">{{.yes}}</label>
				<input type="radio" name="q4" id="q4yes" value="Yes" required>
				<label for="q4no">{{.no}}</label>
				<input type ="radio" name="q4" id="q4no" value="No" required>
			</fieldset>
			<h4>V</h4>
			<fieldset id="q5" form="application">
				<p>{{.q5}}</p>
				<label for="q5yes">{{.yes}}</label>
				<input type="radio" name="q5" id="q5yes" value="Yes" required>
				<label for="q5no">{{.no}}</label>
				<input type ="radio" name="q5" id="q5no" value="No" required>
			</fieldset>
			<h4>VI</h4>
			<fieldset id="q6">
				<p>{{.q6}}</p>
				<label for="q6yes">{{.yes}}</label>
				<input type="radio" name="q6" id="q6yes" value="Yes" required>
				<label for="q6no">{{.no}}</label>
				<input type ="radio" name="q6" id="q6no" value="No" required>
			</fieldset>
			<h4>VII</h4>
			<fieldset id="q7">
				<p>{{.q7}}</p>
				<label for="q7yes">{{.yes}}</label>
				<input type="radio" name="q7" id="q7yes" value="Yes" required>
				<label for="q7no">{{.no}}</label>
				<input type ="radio" name="q7" id="q7no" value="No" required>
			</fieldset>
			<h4>VIII</h4>
			<fieldset id="q8">
				<p>{{.q8}}</p>
				<label for="q8yes">{{.yes}}</label>
				<input type="radio" name="q8" id="q8yes" value="Yes" required>
				<label for="q8no">{{.no}}</label>
				<input type ="radio" name="q8" id="q8no" value="No" required>
			</fieldset>
			<h4>IX</h4>
			<fieldset id="q9">
				<p>{{.q9}}</p>
				<label for="q9yes">{{.yes}}</label>
				<input type="radio" name="q9" id="q9yes" value="Yes" required>
				<label for="q9no">{{.no}}</label>
				<input type ="radio" name="q9" id="q9no" value="No" required>
			</fieldset>
			<h4>X</h4>
			<fieldset id="q10">
				<p>{{.q10}}</p>
				<label for="q10yes">{{.yes}}</label>
				<input type="radio" name="q10" id="q10yes" value="Yes" required>
				<label for="q10no">{{.no}}</label>
				<input type ="radio" name="q10" id="q10no" value="No" required>
			</fieldset>
			<h4>XI</h4>
			<fieldset id="q11">
				<p>{{.q11}}</p>
				<label for="q11yes">{{.yes}}</label>
				<input type="radio" name="q11" id="q11yes" value="Yes" required>
				<label for="q11no">{{.no}}</label>
				<input type ="radio" name="q11" id="q11no" value="No" required>
			</fieldset>
			<h4>{{.confirm}}</h4>
			<fieldset id="q12">
				<p>{{.q12}}</p>
				<label for="confirm">{{.confirm}}</label>
				<input type="checkbox" name="redFlag" id="confirm" required>
				<br><!-- clean up the formatting on some browsers -->
			</fieldset>
			<h4>{{.user}}</h4>
			<fieldset>
				<label for="email">{{.email}}</label>
				<input type="email" id="email" name="email" required>
				<label for="name">{{.fullname}}</label>
				<input id="name" name="name" required>
				<input type="hidden" name="token" value="{{.token}}">
			</fieldset>
			<button type="submit">{{.submit}}</button>
		</form>
	</main>
{{end}}
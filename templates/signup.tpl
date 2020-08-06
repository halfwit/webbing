{{define "content"}}
	<main>
		<h2>{{.mainHeader}}</h2>
		{{template "errors" .errors}}
		<form method="post" action="signup.html">
			<label for="fname">{{.fname}}*</label>
			<input type="text" name="fname" id="fname" placeholder="{{.fnameph}}" required autocomplete="off" autofocus/><br>
			<label for="lname">{{.lname}}*</label>
			<input type="text" name="lname" id="lname" placeholder="{{.lnameph}}" required autocomplete="off"/><br>
			<label for="email">{{.email}}*</label>
			<input type="email" name="email" id="email" placeholder="{{.emailph}}" required autocomplete="off"/><br>
			<label for="pass">{{.pass}}*</label>
			<input type="password" name="pass" id="pass" minlength="8" placeholder="{{.passph}}" required autocomplete="off"/><br>
			<button type="submit" class="button button-block"/>{{.gobutton}}</button>
		</form>
	</main>
{{end}}
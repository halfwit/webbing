{{define "content"}}
	<main>
            <h3>{{if .redirect }}{{.continue}}{{else}}{{.greeting}}{{end}}</h3>
	      {{range .errors}}
	      <p style="color: red" class="errtext">{{.}}</p>
	      {{end}}
	      <form method="post" action="login.html">
			<label for="email">{{.email}}*</label>
			<input type="email" name="email" id="email" required autocomplete="off"/><br>
			<label for="pass">{{.password}}*</label>
			<input type="password" name="pass" id="pass" minlength="8" required autocomplete="off"/><br>
			<p class="forgot"><a href="resetpassword.html">{{.forgotPassword}}</a></p>
			<button type="submit" class="button button-block"/>{{.login}}</button>
          	    </form>
	</main>
{{end}}
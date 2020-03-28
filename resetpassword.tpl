{{define "content"}}
	<main>
            <h3>{{.reset}}</h3>
	      {{template "errors" .errors}}
	      <p>{{.resettext}}
	      <form method="post" action="resetpassword.html">
			<label for="email">{{.email}}*</label>
			<input type="email" name="email" id="email" required autocomplete="off"/><br>
			<button type="submit" class="button button-block"/>{{.sendreset}}</button>
          </form>
	</main>
{{end}}
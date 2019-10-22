{{define "content"}}
	<main>
            <h3>{{.reset}}</h3>
	      {{range .errors}}
	      <p style="color: red" class="errtext">{{.}}</p>
	      {{end}}
	      <form method="post" action="{{.basedir}}/newpassword.html">
			<input type="hidden" name="token" value="{{.token}}"/>
			<label for="password">{{.password}}*</label>
			<input type="password" name="password" id="password" required autocomplete="off"/><br>
			<laber for="reenter">{{.reenter}}*</label>
			<input type="password" name="reenter" id="reenter" required autocomplete="off"/><br>
			<button type="submit" class="button button-block"/>{{.update}}</button>
          	    </form>
	</main>
{{end}}
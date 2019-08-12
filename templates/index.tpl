{{define "content"}}
	  <aside class="tagline">
	    <h2>Quality Healthcare</h2>
	      <p>Your Terms</p>
	      <p>Your Time</p>
	      <p>Your Coin</p>
	  </aside>
	  <main>
	    <section>
	      <h2>{{.whoWeAre}}</h2> 
	      <p>{{.aboutUs}}</p>
            </section>
	    <section>
	      <h2>{{.secondOpinions}}</h2>
	      <p>{{.fromHome}}</p>
	    </section>
	    <section>
	      <h2>{{.anonymity}}</h2>
	      <p>{{.anonText}}</p>
	    </section>
	    <section>
	      <h2>{{.wholeWorld}}</h2>
	      <p>{{.wholeWorldText}}</p>
	    </section>
	    <section>
	      <h2>{{.payment}}</h2>
	      <p>{{.paymentText}}</p>
	    </section>
	  </main>
	  <aside class="showcase">
	    <table>
	    <tr>
	        <th colspan="3">{{.doctorsFrom}}<!-- {{.userCountry}} --></th>
		{{range .doctors}}
	    <tr>
	    <th><img src="images/{{.Image}}" width="155" height="255"/><a href="guest/invoice.html"><h4>{{.Name}}</h4></a>
		    <p>{{.Specialty}}
		    <p>{{.Rate}}</p>
		</th>
		</tr>
		{{end}}
		<th><a href="guest/catalog.html"><h5>{{.seeMore}}</h5></a></th>
	    </table>
	    </section>
	  </aside>
{{end}}
{{define "content"}}
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
{{end}}
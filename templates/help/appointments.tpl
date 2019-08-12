{{define "content"}}
	<main>
	<h1>{{.mainHeader}}</h1>
	<h2>{{.requestHeader}}</h2>
	<h3>{{.statusHeader}}</h3>
	<p>{{.statusBody}}</p>
	<br>
	<h3>{{.scheduleHeader}}</h3>
	<p>{{.scheduleBody}}</p>
	<br>
	<h3>{{.expiresHeader}}</h3>
	<p>{{.expiresBody}}</p>
	<br>
	<h3>{{.emailHeader}}</h3>
	<p>{{.emailBody1}}</p>
	<p>{{.emailBody2}}</p>
    	<p style="margin-right: 40px">{{.emailBody3}}</p>
    	<p style="margin-right: 40px">{{.emailBody4}}</p>
    	<p style="margin-right: 40px">{{.emailBody5}}</p>
	<h4>{{.notifyHeader}}</h4>
	<p>{{.notifyBody}}</p>
    	<p style="margin-right: 40px">{{.notifyBody1}}</p>
    	<p style="margin-right: 40px">{{.notifyBody2}}</p>
    	<p style="margin-right: 40px">{{.notifyBody3}}</p>
	<br>
	<h4>{{.inboxSearchHeader}}</h4>

	<p>{{.inboxSearchBody}}</p>

	<h4>{{.checkSpamHeader}}</h4>
	<p>{{.checkSpamBody}}</p>

	<p style="margin-right: 40px">{{.checkSpamBody1}}</p>
    	<p style="margin-right: 40px">{{.checkSpamBody2}}</p>

	<p>{{.checkSpamAdd}}</p></br>
	<h4>{{.deliveryHeader}}</h4>
	<br>
	<h4>{{.blockHeader}}</h4> 
	<p>{{.blockBody}}</p>
	</main>
{{end}}
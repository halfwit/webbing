{{define "content"}}
	<main>
	    <h1>{{.mainHeader}}</h1>
		<h2>{{.helpHeader}}</h2>
		<h4>{{.topics}}</h4>
		<h4>{{.appointmentHeader}}</h4>
		<p>{{.appointmentStatus}}</p>
		<p>{{.appointmentClear}}</p>
		<p>{{.appointmentExpire}}</p>
		<a href="appointments.html">{{.viewall}}</a>
		<h4>{{.paymentHeader}}</h4>
		<p>{{.paymentEdit}}</p>
		<p>{{.paymentBitcoin}}</p>
		<p>{{.paymentBitcoinHow}}</p>
		<p>{{.paymentAdd}}</p>
		<a href="paymentmethods.html">{{.viewall}}</a>
		<h4>{{.verify}}</h4>
		<p>{{.verifyPhone}}</p>
		<p>{{.verifyEmail}}</p>
		<p>{{.verifyLicense}}</p>
		<a href="verification.html">{{.viewall}}</a>
		<h4>{{.priceHeader}}</h4>
		<p>{{.priceDetermined}}</p>
		<p>{{.priceWhen}}</p>
		<p>{{.priceCurrency}}</p>
		<a href="pricesandfees.html">{{.viewall}}</a>
		<h4>{{.contactHeader}}</h4>
		<p>{{.contactStatus}}</p>
		<p>{{.contactHow}}</p>
		<a href="contacting.html">{{.viewall}}</a>
	</main>
{{end}}
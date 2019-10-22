{{define "content"}}
	<main>
	<h1>{{.mainHeader}}</h1>
	<h2>{{.paymentHeader}}</h2>
		<p>{{.paymentBody}}<a href="https:/bitcoin.org">Bitcoin</a></p>
	<h4>{{.whatBTCHeader}}</h4>
		<p>{{.whatBTCBody}}<a href="https://en.wikipedia.org/wiki/Bitcoin">Bitcoin</a></p>
		<br>
	<h4>{{.chargedHeader}}</h4>
		<p>{{.chargedBody}}</p>
		<br>
	<h4>{{.anyCoinHeader}}</h4>
		<p>{{.anyCoinBody}}</p>	
		<br>	
	<h4>{{.editWalletHeader}}</h4>
		<p>{{.editWalletBody}}</p>
		<br>
	<h4>{{.delWalletHeader}}</h4>
		<p>{{.delWalletBody}}</p>
    	<p>{{.delWalletBody1}}</p>
    	<p>{{.delWalletBody2}}</p>
		<br>
<p>{{.otherCoinNote}}</p>	
		<br>
	<h4>{{.blockedHeader}}</h4> 
<p>{{.blockedBody}}<a href="https://www.torproject.org">Tor</a></p>
	</main>
{{end}}
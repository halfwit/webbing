{{define "header"}}
<!DOCTYPE html>
<html>
<head>
  <title>{{.title}}</title>
  {{if .css}}
  <link rel="stylesheet" text="text/css" href="{{.basedir}}/css/{{.css}}">
  {{end}}
  <link rel="icon" type="image/svg+xml" href="favicon.svg">
  <link rel="stylesheet" test="text/css" href="{{.basedir}}/css/default.css">
  <script type="text/javascript" src="{{.basedir}}/scripts/details-polyfill.js"></script>
</head>
<body>
  <header class="navbar">
    <!-- This changes based on user login status-->
    <a href="{{.basedir}}/index.html" class="logo"><svg class="headerlogo" xmlns="http://www.w3.org/2000/svg" height="380pt" viewBox="0 0 380 380" width="380pt"><path d="m360.296875 172.976562-170.632813-170.632812c-.734374-.742188-1.621093-1.328125-2.613281-1.734375-1.953125-.808594-4.152343-.808594-6.113281 0-.984375.40625-1.871094 1-2.609375 1.734375l-170.621094 170.632812-5.359375 5.359376c-3.128906 3.128906-3.128906 8.183593 0 11.3125l175.992188 175.992187c.734375.742187 1.621094 1.328125 2.605468 1.734375.976563.410156 2.015626.617188 3.058594.617188 1.039063 0 2.078125-.207032 3.054688-.617188.984375-.40625 1.871094-.992188 2.605468-1.734375l175.992188-175.992187c3.128906-3.128907 3.128906-8.183594 0-11.3125zm-184.292969 48.777344-151.089844-37.761718 151.089844-37.761719zm16-75.523437 151.085938 37.761719-151.085938 37.761718zm-16-16.476563-136.570312 34.125 136.570312-136.574218zm0 108.476563v102.449219l-136.570312-136.574219zm16 0 136.566406-34.125-136.566406 136.574219zm0-108.476563v-102.449218l136.566406 136.566406zm0 0"/></svg> Olmax Medical</a>
    <input type="checkbox" id="showmenu" />
    <label for="showmenu">&#x2630;</label>
    <nav class="hamburger">
      <details class="doctors">
        <summary>{{.doctors}}</summary>
        <ul>
          <li><a href="{{.basedir}}/guest/becomeAProvider.html">{{.provider}}</a></li>
	  <li><a href="{{.basedir}}/guest/whocandoctor.html">{{.whodoctor}}</a></li>
        </ul>
      </details>
      <details class="how-it-works">
        <summary>{{.howworks}}</summary>
        <ul>
          <li><a href="{{.basedir}}/help/appointments.html">{{.appts}}</a></li>
          <li><a href="{{.basedir}}/help/contacting.html">{{.proc}}</a></li>
          <li><a href="{{.basedir}}/help/paymethods.html">{{.payments}}</a></li>
	  <li><a href="{{.basedir}}/help/pricefee.html">{{.fees}}</a></li>
          <li><a href="{{.basedir}}/help/paymethods.html">{{.verify}}</a></li>
        </ul>
      </details>
      <ul>
        <li><a href="{{.basedir}}/help/faq.html">{{.faq}}</a></li>
        {{if eq .status "true"}}
	<li><a href="{{.basedir}}/profile.html">{{.profile}}</a></li>
        {{end}}
      </ul>
      <hr/>
      <details class="contact-us">
        <summary>{{.contact}}</summary>
        <ul>
	  <li><a href="mailto:help@olmaxmedical.com">{{.email}}</a></li>
 	  <li><a href="tel:1-555-555-5555">{{.phone}}</a></li>
        </ul>
      </details>
    </nav>
    <nav class="login">
      <ul>
	{{if eq .status "true"}}
	<li><a class="signin" href="{{.basedir}}/logout.html">{{.logout}}</a></li>
	{{else}}
        <li><a class="signup" href="{{.basedir}}/signup.html">{{.signup}}</a></li>
        <li><a class="signin" href="{{.basedir}}/login.html">{{.login}}</a></li>
	{{end}}
      </ul>
    </nav>
  </header>
{{end}}
{{define "header"}}
<!DOCTYPE html>
<html>
<head>
  <title>{{.title}}</title>
  {{if .css}}
  <link rel="stylesheet" text="text/css" href="{{.basedir}}/css/{{.css}}">
  {{end}}
  <link rel="icon" type="image/svg+xml" href="{{.basedir}}/favicon.svg">
  <link rel="stylesheet" test="text/css" href="{{.basedir}}/css/default.css">

</head>
<body>
  <header class="navbar">
    <!-- This changes based on user login status-->
    <nav>
      <ul>
        <li><a href="{{.basedir}}/index.html">{{.home}}</a></li>
  {{if eq .status "true"}}
	      <li><a href="{{.basedir}}/profile.html">{{.profile}}</a></li>
	      <li><a href="{{.basedir}}/logout.html">{{.logout}}</a></li>
	{{else}}
        <li><a href="{{.basedir}}/signup.html">{{.signup}}</a></li>
        <li><a href="{{.basedir}}/login.html">{{.login}}</a></li>
	{{end}}
      </ul>
    </nav>
  </header>
{{end}}
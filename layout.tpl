{{define "layout"}}
{{template "header" .header}}
	<div class="content">
{{template "content" .}}
	</div>
{{template "footer" .footer}}
{{end}}
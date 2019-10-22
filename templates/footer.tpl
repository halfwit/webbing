{{define "footer"}}
 <footer>
    <section class="branding">
      <a href="{{.basedir}}/index.html"><svg xmlns="http://www.w3.org/2000/svg" height="380pt" viewBox="0 0 380 380" width="380pt"><path d="m360.296875 172.976562-170.632813-170.632812c-.734374-.742188-1.621093-1.328125-2.613281-1.734375-1.953125-.808594-4.152343-.808594-6.113281 0-.984375.40625-1.871094 1-2.609375 1.734375l-170.621094 170.632812-5.359375 5.359376c-3.128906 3.128906-3.128906 8.183593 0 11.3125l175.992188 175.992187c.734375.742187 1.621094 1.328125 2.605468 1.734375.976563.410156 2.015626.617188 3.058594.617188 1.039063 0 2.078125-.207032 3.054688-.617188.984375-.40625 1.871094-.992188 2.605468-1.734375l175.992188-175.992187c3.128906-3.128907 3.128906-8.183594 0-11.3125zm-184.292969 48.777344-151.089844-37.761718 151.089844-37.761719zm16-75.523437 151.085938 37.761719-151.085938 37.761718zm-16-16.476563-136.570312 34.125 136.570312-136.574218zm0 108.476563v102.449219l-136.570312-136.574219zm16 0 136.566406-34.125-136.566406 136.574219zm0-108.476563v-102.449218l136.566406 136.566406zm0 0"/></svg></a>
      <h1>Olmax Medical</h1>
    	<banner>{{.banner}}</banner>
    	<small>{{.copy}}</small>
    </section>
    <nav class="leftfooter">
      <details>
	<summary>{{.aboutHeader}}</summary>
	<ul>
          <li><a href="{{.basedir}}/aboutus.html">{{.about}}</a></li>
	</ul>
      </details>
      <details class="partner">
        <summary>{{.partHead}}</summary>
	<ul>
          <li><a href="{{.basedir}}/partner.html">{{.partner}}</a></li>
          <li><a href="{{.basedir}}/becomeAProvider.html">{{.provider}}</a></li>
	</ul>
      </details>
    </nav>
    <nav class="rightfooter">
      <details class="legal">
        <summary>{{.legal}}</summary>
	<ul>
          <li><a href="{{.basedir}}/help/legal.html">{{.legal}}</a></li>
	  <li><a href="{{.basedir}}/help/privacy.html">{{.privacy}}</a></li>
	</ul>
      </details>
    </nav>
  </footer>
</body>

</html>
{{end}}
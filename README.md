# Creating Pages

Starting from html, a page is translated into a template. A template is the html with the header/footer removed. Example:

```html
<div class="content">
	<main>
		<section>
			<h1>Something</h1>
			<p>Content</p>
		</section>
	</main>
	<aside>
		<h2>Something else</h2>
		<p>Content</p>
	</aside>
</div>
```

Then, each `string` is taken out so that it can be ran through the internationalization stuff

```html
<div class="content">
	<main>
		<section>
			<h1>{{.mainHeader}}</h1>
			<p>{{.mainContent}}</p>
		</section>
	</main>
	<aside>
		<h2>{{.otherHeader}}</h2>
		<p>{{.otherContent}}</p>
	</aside>
</div>
```

The {{.foo}} corresponds to the contents of pages/path/to/this/page.go. A map[string]interface{} is used, with an entry for each key above, and the value.

```go
import (
	"golang.org/x/text/message"
)

func Mypage(p *message.Printer) map[string]interface{} {
	return map[string]interface{} {
		"mainHeader": p.Sprint("Something"),
		"mainContent": p.Sprint("Content"),
		"otherHeader": p.Sprint(Something else"),
		"otherContent: p.Sprint("Content"),
	}
}
```

Now, the program when it starts up will run the init() function ofanything imported. In client.go, you would import the library that this page you've made belongs to as an empty import - this basically means it'll always run the init(). We use this to our advantage

```go
func init() {
	b := &router.Page{
		// Who can access the page
		Access: router.GuestAuth|router.PatientAuth, // or router.DoctorAuth
		// There's a basic CSS for every given page that is already applied
		// we use this if we want special css for this page
		Css: "customstuffforonlythispage.css",
		// would relate to, for example, olmaxmedical.com/path/to/this/page.html
		Path: path/to/this/page 
		Data: Mypage, // pointer to function we defined above
		// These load your interface{} item above so you can access the lists
		// This keeps the API boundaries clean, and more importantly drastically
		// simplifies writing more pages
		Extra: router.ListCountries|router.ListDoctors|router.ListSpecialties, 
}
```

An example using one of the above extras would be like so

```html
<div class="content">
	<main>
		<section>
			<h1>Something</h1>
			<p>Content</p>
		</section>
	</main>
	<aside>
		<form action="dosomething">
			<select name="Specialties">
			{{range .specialties}} <!-- this is from ListSpecialties-->
				<option value="{{.ID}}">{{.Name}}</option>
			{{end}}
			</select>
		</form>
		<h2>Something else</h2>
		<p>Content</p>
	</aside>
</div>
```

The structs for all of these are listed in their source files in router/, as they're only ever loaded and used there.

So each page is basically `cached` at startup, and calling them has only the overhead of translation, which is very quick map lookups based on a numerical identifier.
There's no further steps required; just make the two files, possibly the extra CSS, and if you are creating a new library (that is to say, a path in olmaxmedical.com/some/new/area.html) make sure it gets added in client.go

```html
import (
	...
	_ "olmax/some/new"
	...
)
```

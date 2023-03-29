package main

import "fmt"
import "html/template"
import "net/http"

const view string = `<html>
<head>
	<title>Template Render HTML String</title>
	</head>
	<body>
		<h1>Template Render HTML String</h1>
		<p>Template Render HTML String</p>
		</body>
		</html>`

func main() {
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		var tmpl = template.Must(template.New("main-template").Parse(view))
		if err := tmpl.Execute(w, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/index", http.StatusTemporaryRedirect)
	})

	fmt.Println("starting web server at http://localhost:9000/")
	http.ListenAndServe(":9000", nil)
}
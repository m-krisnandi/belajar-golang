package main

import "net/http"
import "html/template"
import "fmt"

type Superhero struct {
	Name string
	Alias string
	Friends []string
}

func (s Superhero) SayHello(from string, message string) string {
	return fmt.Sprintf("%s said: \"%s\"", from, message)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var person = Superhero {
			Name: "Clark Kent",
			Alias: "Superman",
			Friends: []string{"Batman", "Flash", "Green Lantern"},
		}

		var tmpl = template.Must(template.ParseFiles("view.html"))
		if err := tmpl.Execute(w, person); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("Starting web server at http://localhost:9000/")
	http.ListenAndServe(":9000", nil)
}
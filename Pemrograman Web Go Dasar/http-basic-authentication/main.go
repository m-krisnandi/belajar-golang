package main

import (
	"fmt"
	"encoding/json"
	"net/http"
)

// testing
// go run *.go
// $ curl -X GET --user admin:secret http://localhost:9000/student
// $ curl -X GET --user admin:secret http://localhost:9000/student?id=s001

func main() {
	http.HandleFunc("/student", ActionStudent)

	server := new(http.Server)
	server.Addr = ":9000"

	fmt.Println("Server started at localhost:9000")
	server.ListenAndServe()
}

func ActionStudent(w http.ResponseWriter, r *http.Request) {
	if !Auth(w, r) { return }
	if !AllowOnlyGET(w, r) { return }

	if id := r.URL.Query().Get("id"); id != "" {
		OutputJSON(w, SelectStudent(id))
		return
	}

	OutputJSON(w, GetStudents())
}

// konversi data menjadi JSON string
func OutputJSON(w http.ResponseWriter, o interface{}) {
	res, err := json.Marshal(o)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
package main

import "fmt"
import "net/http"
import "encoding/json"

func main() {
	http.HandleFunc("/", ActionIndex)

	fmt.Println("starting web server at http://localhost:9000/")
	http.ListenAndServe(":9000", nil)
}

func ActionIndex(w http.ResponseWriter, r *http.Request) {
	data := []struct {
		Name string
		Age int
	} {
		{"Richard Hendricks", 24},
		{"Erlich Bachman", 30},
		{"Bertram Gilfoyle", 30},
		{"Dinesh Chugtai", 28},
	}

	// cara 1
	// jsonInBytes, err := json.Marshal(data)

	// cara 2
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	// w.Write(jsonInBytes)

	
}
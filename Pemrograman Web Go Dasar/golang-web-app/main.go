package main

import "fmt"
import "net/http"

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	var message = "Welcome"
	w.Write([]byte(message))
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	var message = "Hello World"
	w.Write([]byte(message))
}

func main() {
	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/index", handlerIndex)
	http.HandleFunc("/hello", handleHello)

	var address = "localhost:9000"
	fmt.Printf("server started at %s\n", address)
	// cara 1
	// err := http.ListenAndServe(address, nil)

	// cara 2
	server := new(http.Server)
	server.Addr = address
	err := server.ListenAndServe()

	// timeout untuk read dan write request
	// server.ReadTimeout = time.Second * 10
	// server.WriteTimeout = time.Second * 10

	if err != nil {
		fmt.Println(err.Error())
	}
}
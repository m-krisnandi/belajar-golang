package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

// testing with curl
// curl -X GET http://localhost:9000/

// testing with curl with payload
// curl -X POST http://localhost:9000/ -H 'Content-Type: application/json' -d '{}'


func handleIndex(w http.ResponseWriter, r *http.Request) {
	done := make(chan bool)
	go func() {
		// do the process here
		// simluate a long-time request by putting 10 seconds sleep
		// time.Sleep(10 * time.Second)

		// // send a signal to the channel
		// done <- true


		// Handle Cancelled Request yang ada Payload-nya
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Error reading body:", err.Error())
			return
		}
		log.Println("Body:", string(body))

		// simluate a long-time request by putting 10 seconds sleep
		time.Sleep(10 * time.Second)

		// send a signal to the channel
		done <- true
	}()

	select {
	case <-r.Context().Done():
		if err := r.Context().Err(); err != nil {
			if strings.Contains(strings.ToLower(err.Error()), "canceled") {
				log.Println("Request canceled by the client")
			} else {
				log.Println("Unknown error:", err.Error())
			}
		}
	case <-done:
		log.Println("Request processed")
	}
}

func main() {
	http.HandleFunc("/", handleIndex)
	http.ListenAndServe(":9000", nil)
}
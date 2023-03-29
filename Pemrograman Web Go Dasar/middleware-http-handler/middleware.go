package main

import "net/http"

const USERNAME = "admin"
const PASSWORD = "secret"

func MiddlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			w.Write([]byte("something wrong with basic auth"))
			return 
		}

		isValid := (username == USERNAME) && (password == PASSWORD)
		if !isValid {
			w.Write([]byte("invalid username or password"))
			return 
		}

		next.ServeHTTP(w, r)
	})
}

func MiddlewareAllowOnlyGET(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.Write([]byte("only GET method allowed"))
			return 
		}
	
		next.ServeHTTP(w, r)
	})
	
}
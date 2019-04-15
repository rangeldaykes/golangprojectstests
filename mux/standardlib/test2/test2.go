package main

import (
	"context"
	"fmt"
	"net/http"
)

func main() {
	dashboard := http.NewServeMux()
	dashboard.HandleFunc("/dashboard/hi", printHi)
	dashboard.HandleFunc("/dashboard/bye", printBye)

	mux := http.NewServeMux()
	mux.Handle("/dashboard/", requireUser(dashboard))
	mux.HandleFunc("/", home)

	http.ListenAndServe(":3000", addRequestID(mux))
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Homepage...")
}

func printHi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hi! Your request ID is:", r.Context().Value("request_id"))
}

func printBye(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Bye! Your request ID is:", r.Context().Value("request_id"))
}

var requestID = 0

func nextRequestID() int {
	requestID++
	return requestID
}

func addRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "request_id", nextRequestID())
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

type User struct{}

func lookupUser(r *http.Request) *User {
	return &User{}
}

func requireUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := lookupUser(r)
		if user == nil {
			// No user so redirect to login
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}

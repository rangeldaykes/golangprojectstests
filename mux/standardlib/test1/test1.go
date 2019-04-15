package main

import (
	"fmt"
	"net/http"
)

type User struct {
}

func lookupUser(r *http.Request) *User {
	return &User{}
}

func requireUser(fn func(http.ResponseWriter, *http.Request, *User)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := lookupUser(r)
		if user == nil {
			// No user so redirect to login
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		fn(w, r, user)
	}
}

func printUser(w http.ResponseWriter, r *http.Request, user *User) {
	fmt.Fprintln(w, "User is:", user)
}

func main() {
	http.HandleFunc("/user", requireUser(printUser))
	http.ListenAndServe(":3000", nil)
}

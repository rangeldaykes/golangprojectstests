// https://www.jtolio.com/2017/01/writing-advanced-web-applications-with-go/
package main

import "net/http"

func RequireUser(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		user, err := GetUser(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if user == nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		h.ServeHTTP(w, req)
	})
}

func main() {
	// v1 http.ListenAndServe(":8080", RequireUser(http.HandlerFunc(myHandler)))

	// v2
	mux := http.NewServeMux()
	mux.Handle("/user/", RequireUser(http.HandlerFunc(myHandler)))
	http.ListenAndServe(":8080", mux)
}

type User struct {
	ID   int
	Name string
}

func GetUser(req *http.Request) (*User, error) {
	return &User{ID: 1, Name: "zé lelé"}, nil
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("return user"))
}

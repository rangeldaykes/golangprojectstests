package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", StatusPage)
	mux.HandleFunc("/login", LoginPage)
	mux.HandleFunc("/logout", LogoutPage)

	contexttedMux := AddContext(mux)

	log.Println("Start server on port :8085")

	//log.Fatal(http.ListenAndServe(":8085", mux))
	log.Fatal(http.ListenAndServe(":8085", contexttedMux))
}

func StatusPage(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("This Page will show the context username once the context is added."))

	// get data from context
	if username := r.Context().Value("Username"); username != nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello " + username.(string) + "\n"))
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Logged in"))
	}
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	expiration := time.Now().Add(365 * 24 * time.Hour) // set to expire in 1 year
	cookie := http.Cookie{Name: "username", Value: "name@gmail.com", Expires: expiration}
	http.SetCookie(w, &cookie)
}

func LogoutPage(w http.ResponseWriter, r *http.Request) {
	expiration := time.Now().AddDate(0, 0, -1) // set to expire in the past
	cookie := http.Cookie{Name: "username", Value: "name@name.com", Expires: expiration}
	http.SetCookie(w, &cookie)
}

func AddContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, "-", r.RequestURI)
		cookie, _ := r.Cookie("username")
		if cookie != nil {
			// add data to context
			ctx := context.WithValue(r.Context(), "Username", cookie.Value)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			next.ServeHTTP(w, r)
		}

	})
}

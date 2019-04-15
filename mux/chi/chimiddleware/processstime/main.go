package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/middleware"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(ProcessingTime)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("process time", "aaa")
		w.Write([]byte("hello world"))
	})

	http.ListenAndServe(":8000", r)
}

// ProcessingTime is a mid
func ProcessingTime(next http.Handler) http.Handler {

	fn := func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()

		w.Header().Set("kkk", "lll")

		next.ServeHTTP(w, r)

		t2 := time.Now()
		diff := t2.Sub(t1)
		diffmili := int64(diff / time.Millisecond)

		//w.Header().Set("X-Processing-Time", string(strconv.FormatInt(diffmili, 10)))

		w.Header().Set("mmm", "nnn")

		sdiff := strconv.FormatInt(diffmili, 10)

		w.Write([]byte(sdiff))
	}

	return http.HandlerFunc(fn)
}

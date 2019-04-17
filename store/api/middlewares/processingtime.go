package middlewares

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// ProcessingTime is a mid
func ProcessingTime(next http.Handler) http.Handler {

	fn := func(w http.ResponseWriter, r *http.Request) {

		t1 := time.Now()
		//headerori := &w
		//w.Header().Write()

		w.Header().Set("kkk", "lll")

		next.ServeHTTP(w, r)

		//next.ServeHTTP(*headerori, r)

		t2 := time.Now()
		diff := t2.Sub(t1)
		diffmili := int64(diff / time.Millisecond)
		fmt.Println(diffmili)

		w.Header().Set("X-Processing-Time", string(strconv.FormatInt(diffmili, 10)))

		//for a, b := range w.Header() {
		//fmt.Println(a, b)
		//}

		w.Header().Set("mmm", "nnn")

		//for a, b := range w.Header() {
		//fmt.Println(a, b)
		//}

		//(*headerori).Header().Set("mmm", "nnn")

		//w.Write([]byte("tgtg"))

	}

	return http.HandlerFunc(fn)
}

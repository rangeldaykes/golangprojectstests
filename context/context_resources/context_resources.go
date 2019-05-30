package context_resources

import (
	"context"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Resources)

	log.Println("Start server on port :8085")
	log.Fatal(http.ListenAndServe(":8085", mux))
}

func Resources(w http.ResponseWriter, r *http.Request) {
	// Timeout in context
	context.WithTimeout(
		r.Context(),
		time.Duration(1000)*time.Millisecond,
	)
}

type Service struct {
	ctx context.Context
}

func New(ctx context.Context) *Service {
	return &Service{ctx}
}

func (s Service) Get([]byte) {

}

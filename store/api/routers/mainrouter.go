package routers

import (
	"net/http"
	"store/api/middlewares"
	"store/infrastucture/injectioncontainer"
	"sync"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// IChiRouter Initialize Router
type IChiRouter interface {
	InitRouter() *chi.Mux
}

type router struct{}

var (
	m          *router
	routerOnce sync.Once
)

// ChiRouterGetInstance get a IChiRouter
func ChiRouterGetInstance() IChiRouter {
	if m == nil {
		routerOnce.Do(func() {
			m = &router{}
		})
	}

	return m
}

func (router *router) InitRouter() *chi.Mux {

	r := chi.NewRouter()

	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	//r.Use(middlewares.ProcessingTime)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("home"))
	})

	r.Mount("/api/v1/User", router.RoutesUser())

	return r
}

func (router *router) RoutesUser() chi.Router {
	r := chi.NewRouter()

	r.Use()

	// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	injectioncontainer.InjectUserController().List(w, r)
	// })

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			injectioncontainer.InjectUserController().Get(w, r)
		})
	})

	return r
}


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

package routers

import (
	"net/http"
	"store/infra/di"

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
	r.Use(MiddlewareBasic)

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			//rootHandler(di.InjectUserController().Get).ServeHTTP(w, r)
			di.InjectUserController().Get(w, r)
		})
	})

	return r
}

type rootHandler func(http.ResponseWriter, *http.Request) error

func (fn rootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := fn(w, r) // Call handler function.

	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func MiddlewareBasic(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

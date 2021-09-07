package main

import (
	"log"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/DooomiT/suff.de-v-backend/pkg/api/alcohol"
	"github.com/DooomiT/suff.de-v-backend/pkg/api/snack"
	"github.com/DooomiT/suff.de-v-backend/pkg/api/softdrink"

)

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/api/alcohol", alcohol.Routes())
		r.Mount("/api/snack", snack.Routes())
		r.Mount("/api/softdrink", softdrink.Routes())

	})

	return router
}
func main() {
	router := Routes()

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route) // Walk and print out all routes
		return nil
	}
	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging err: %s\n", err.Error()) // panic if there is an error
	}

	log.Fatal(http.ListenAndServe(":8000", router)) // Note, the port is usually gotten from the environment.
}

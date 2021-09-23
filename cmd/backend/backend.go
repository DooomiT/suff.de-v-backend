<<<<<<< HEAD
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
=======
package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/DooomiT/suff.de-v-backend/pkg/db"
	"github.com/DooomiT/suff.de-v-backend/pkg/handler"
)

const (
	USERNAME = "postgres"
	PASSWORD = "2002"
	HOST     = "localhost"
	DBNAME   = "testdb"
	PORT     = 5432
)

const local = true

func getDatabaseEnv() (string, string, string, string, string) {
	return os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT")
}

func main() {
	addr := ":8000"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Error occurred: %s", err.Error())
	}
	var dbUser, dbPassword, dbName, dbHost string
	var dbPort int
	if local {
		dbUser, dbPassword, dbHost, dbName, dbPort = USERNAME, PASSWORD, HOST, DBNAME, PORT
	} else {
		dbUser_env, dbPassword_env, dbName_env, dbHost_env, dbPort_env := getDatabaseEnv()
		dbUser, dbPassword, dbName, dbHost = dbUser_env, dbPassword_env, dbName_env, dbHost_env
		dbPort, err = strconv.Atoi(dbPort_env)
		if err != nil {
			log.Fatalf("Could not convert POSTGRES_PORT to int: %v", err)
		}
	}
	database, err := db.Initialize(dbUser, dbPassword, dbName, dbHost, dbPort)
	if err != nil {
		log.Fatalf("Could not set up database: %v", err)
	}
	defer database.Conn.Close()

	httpHandler := handler.NewHandler(database)
	server := &http.Server{
		Handler: httpHandler,
	}
	go func() {
		server.Serve(listener)
	}()
	defer Stop(server)
	log.Printf("Started server on %s", addr)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(fmt.Sprint(<-ch))
	log.Println("Stopping API server.")
}
func Stop(server *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Could not shut down server correctly: %v\n", err)
		os.Exit(1)
	}
}
>>>>>>> main

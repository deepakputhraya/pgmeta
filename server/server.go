package main

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"pgmeta/server/routes"
	"time"
)

func getEnvironment(key string, defaultValue string) string {
	val := os.Getenv(key)
	if len(val) == 0 {
		return defaultValue
	}
	return val
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func main() {
	url := getEnvironment("PG_META_DB_URL", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	db, err := sqlx.Connect("postgres", url)
	log.Println("Server connecting to ", url)
	if err != nil {
		log.Fatalln(err)
	}
	db = db.Unsafe()

	r := mux.NewRouter()

	// Middlewares
	r.Use(loggingMiddleware)
	// Routes
	routes.SchemaRouter(db, r)
	routes.TableRouter(db, r.PathPrefix("/schema/{schema}").Subrouter())
	routes.ColumnsRouter(db, r.PathPrefix("/schema/{schema}/table/{table}").Subrouter())

	// Handler
	http.Handle("/", r)

	// Print all routes
	_ = r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		m, err := route.GetMethods()
		log.Println(m, t)
		return nil
	})

	// Cors
	handler := cors.Default().Handler(r)

	srv := &http.Server{
		Handler: handler,
		Addr:    ":8080",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}
	log.Println(srv)
	log.Fatal(srv.ListenAndServe())
}

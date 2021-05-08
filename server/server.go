package main

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
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

func main() {
	url := getEnvironment("PG_META_DB_URL", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	db, err := sqlx.Connect("postgres", url)
	log.Println("Server connecting to ", url)
	if err != nil {
		log.Fatalln(err)
	}
	db = db.Unsafe()

	r := mux.NewRouter()
	routes.SchemaRouter(db, r)
	routes.TableRouter(db, r.PathPrefix("/schema/{schema}").Subrouter())
	routes.ColumnsRouter(db, r.PathPrefix("/schema/{schema}/table/{table}").Subrouter())

	http.Handle("/", r)

	_ = r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		m, err := route.GetMethods()
		log.Println(m, t)
		return nil
	})

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	log.Println(srv)
	log.Fatal(srv.ListenAndServe())
}

package routes

import (
	"encoding/json"
	"github.com/deepakputhraya/pgmeta"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
)

func SchemaRouter(db *sqlx.DB, r *mux.Router) {
	// List all schemas
	r.Path("/schemas").
		Methods("GET").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			schemas, err := pgmeta.ListSchemas(db)
			if err != nil {
				log.Println(err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			w.Header().Add("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(schemas)
		})
	subRoute := r.PathPrefix("/schema").Subrouter()
	// Get Schema
	subRoute.Path("/{schema}").
		Methods("GET").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			schema, err := pgmeta.GetSchema(db, vars["schema"])
			if err != nil {
				log.Println(err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			w.Header().Add("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(schema)
		})
	// Create Schema
	subRoute.
		Methods("POST").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "Not found", http.StatusNotFound)
		})
}

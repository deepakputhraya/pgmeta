package routes

import (
	"encoding/json"
	"github.com/deepakputhraya/pgmeta"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"net/http"
)

func TableRouter(db *sqlx.DB, r *mux.Router) {
	// Get all tables in a schema
	r.PathPrefix("/tables").
		Methods("GET").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			tables, err := pgmeta.ListTables(db, vars["schema"])
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
			w.Header().Add("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(tables)
		})
}

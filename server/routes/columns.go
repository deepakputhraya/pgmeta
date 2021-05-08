package routes

import (
	"encoding/json"
	"github.com/deepakputhraya/pgmeta"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"net/http"
)

func ColumnsRouter(db *sqlx.DB, r *mux.Router) {
	// Get all columns
	r.Path("/columns").
		Methods("GET").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			columns, err := pgmeta.ListColumns(db, vars["schema"], vars["table"])
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
			w.Header().Add("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(columns)
		})

	r.Path("/column").
		Methods("POST").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "Not found", http.StatusNotFound)
		})
}

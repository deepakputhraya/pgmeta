package routes

import (
	"encoding/json"
	"github.com/deepakputhraya/pgmeta"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"strings"
)

func TableRouter(db *sqlx.DB, r *mux.Router) {
	// Get all tables in a schema
	r.PathPrefix("/tables").
		Methods("GET").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			tables, err := pgmeta.ListTables(db, vars["schema"])
			if err != nil {
				log.Println(err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
			w.Header().Add("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(tables)
		})

	// Get a table
	r.HandleFunc("/table/{table}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		table, err := pgmeta.GetTable(db, vars["schema"], vars["table"])
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(table)
	}).Methods("GET")

	// Delete or Truncate a table
	r.HandleFunc("/table/{table}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		truncate := false
		if strings.TrimSpace(strings.ToLower(r.URL.Query().Get("truncate"))) == "true" {
			truncate = true
		}

		err := pgmeta.DeleteTable(db, vars["schema"], vars["table"], !truncate)
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"success": true,
		})
	}).Methods("DELETE")

}

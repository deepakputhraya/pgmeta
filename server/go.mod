module pgmeta/server

go 1.15

replace github.com/deepakputhraya/pgmeta => ../

require (
	github.com/deepakputhraya/pgmeta v0.0.0-00010101000000-000000000000
	github.com/gorilla/handlers v1.5.1 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/jmoiron/sqlx v1.3.3
	github.com/lib/pq v1.3.0
	github.com/rs/cors v1.7.0
)

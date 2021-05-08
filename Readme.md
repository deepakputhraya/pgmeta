# pgmeta

A simple library for managing your Postgres. Fetch tables, add roles, and run queries (and more).
This is inspired by [postgres-meta](https://github.com/supabase/postgres-meta)

This is still a WIP progress. Pull Requests are welcome!

## Quickstart

### Install

```shell script
go get github.com/deepakputhraya/pgmeta
```

### Server
```shell script
# First export the connection string to the database
export PG_META_DB_URL=postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable

```
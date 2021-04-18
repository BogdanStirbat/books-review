package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	_ "github.com/lib/pq"
)

func handler(schema graphql.Schema) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		result := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: string(query),
		})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}

var db *sql.DB

func main() {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    QueryType,
		Mutation: MutationType,
	})
	if err != nil {
		log.Fatal(err)
	}
	connStr := "host=localhost port=5432 user=book_db_user password=abc123def dbname=book_db sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	//db, err = sql.Open("postgres", "postgres://book_db_user:abc123def@localhost:5432/book_db")
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/graphql", handler(schema))
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

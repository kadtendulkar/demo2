package main

import (
	"net/http"
	mutation "server/mutation"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

var SignUpMemberSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    mutation.Query,
	Mutation: mutation.MutationType,
})

func main() {

	// graphqlHttpHandler := handler.New(&handler.Config{
	// 	Schema:   &SignUpMemberSchema,
	// 	Pretty:   true,
	// 	GraphiQL: true,
	// })
	// http.Handle("/graphql", graphqlHttpHandler)
	// http.ListenAndServe(":8080", nil)

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {

		setupCorsResponse(&w, r)
		if (*r).Method == "OPTIONS" {
			return
		}

		graphqlHttpHandler := handler.New(&handler.Config{
			Schema:   &SignUpMemberSchema,
			Pretty:   true,
			GraphiQL: true,
		})
		graphqlHttpHandler.ServeHTTP(w, r)

	})
	http.ListenAndServe(":8080", nil)
}

func setupCorsResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
}

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// allow cross domain AJAX requests
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		next.ServeHTTP(w, r)
	})
}

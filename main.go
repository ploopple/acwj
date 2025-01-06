package main

import (
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

// Define the schema
var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
    Query: graphql.NewObject(graphql.ObjectConfig{
        Name: "RootQuery",
        Fields: graphql.Fields{
            "hello": &graphql.Field{
                Type: graphql.String,
                Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                    return "Hello, world!", nil
                },
            },
        },
    }),
})

func main() {
    // Create a new GraphQL handler
    h := handler.New(&handler.Config{
        Schema: &schema,
        Pretty: true,
        GraphiQL: true, // Enable GraphiQL, a GraphQL GUI for testing
    })

    // Set up the HTTP server
    http.Handle("/graphql", h)
    http.ListenAndServe(":8080", nil)
}

package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/liquedgit/tokoMeLia/Database"
	"github.com/liquedgit/tokoMeLia/Directives"
	"github.com/liquedgit/tokoMeLia/graph"
	"github.com/liquedgit/tokoMeLia/helper"
	"github.com/liquedgit/tokoMeLia/middlewares"
	"github.com/rs/cors"
	"log"
	"net/http"
)

const defaultPort = "8080"

func main() {
	port := helper.GoDotEnvVariables("PORT")
	if port == "" {
		port = defaultPort
	}

	Database.MigrateTable()

	router := chi.NewRouter()
	router.Use(middlewares.AuthMiddleware)
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{helper.GoDotEnvVariables("ALLOWED_ORIGINS")},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	c := graph.Config{Resolvers: &graph.Resolver{
		DB: Database.GetInstance(),
	}}

	c.Directives.Auth = Directives.Auth

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(c))
	srv.AddTransport(&transport.Websocket{})

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

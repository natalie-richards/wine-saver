package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
	"github.com/natalie-richards/wine-app/cmd/saver/app"
	"github.com/natalie-richards/wine-app/graph"
	"github.com/rs/cors"
)

// This is the main entry point for the server.

const port = "8020"

func main() {
	fmt.Println("Starting")
	a := app.App{}
	// Ensure the app is initialized and passed in as a Resolver dependency
	a.Init()

	err := a.DBConn.Ping(context.Background())
	if err != nil {
		log.Fatal("db connection failed")
	}
	fmt.Println("DB Connected")

	router := chi.NewRouter()
	// Add CORS middleware around every request
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "http://localhost:" + port},
		AllowCredentials: true,
	}).Handler)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		App: &a,
	}}))
	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return r.Host == "localhost:"+port
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	})

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		panic(err)
	}

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

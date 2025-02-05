package main

import (
	"context"
	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"privacy-ex/pkg/ent"
	"privacy-ex/pkg/graph/resolver"
	"strconv"
)

func main() {
	ctx := context.Background()

	client, err := ent.Open(
		"sqlite3",
		"file:ent?mode=memory&cache=shared&_fk=1",
	)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	port := "8081"

	if _, err := strconv.Atoi(port); err != nil {
		log.Fatalf("Invalid port number: %v", err)
	}

	server := handler.NewDefaultServer(resolver.NewSchema(client))
	server.Use(entgql.Transactioner{TxOpener: client})

	http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))

	log.Printf("Connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

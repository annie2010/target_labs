package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	gen "github.com/Pallinder/go-randomdata"
	"github.com/gopherland/target_labs/gql/internal"
	"github.com/gopherland/target_labs/gql/internal/generated"
	"github.com/gopherland/target_labs/gql/internal/generated/model"
)

const defaultPort = "5000"

//go:generate go run github.com/99designs/gqlgen
func main() {
	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: seedResolver(),
			},
		),
	)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("[BookQL] listening on %s", defaultPort)
	log.Fatal(http.ListenAndServe(":"+defaultPort, nil))
}

func seedResolver() *internal.Resolver {
	r := internal.NewResolver()
	for i := 0; i < 10; i++ {
		isan := fmt.Sprintf("isan-%d", i)
		r.Authors[isan] = model.Author{
			ID:    isan,
			First: gen.FirstName(gen.RandomGender),
			Last:  gen.LastName(),
		}
	}

	victims := make([]model.Author, 0, len(r.Authors))
	for _, a := range r.Authors {
		victims = append(victims, a)
	}

	for i := 0; i < 10; i++ {
		isbn := fmt.Sprintf("isbn-%d", i)
		r.Books[isbn] = model.Book{
			ID:       isbn,
			Title:    gen.SillyName(),
			Category: model.AllBookCategory[rand.Intn(len(model.AllBookCategory))],
			Authors: []model.Author{
				victims[rand.Intn(len(victims))],
			},
		}
	}

	return r
}

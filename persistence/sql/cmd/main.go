// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

package main

import (
	"context"
	"database/sql"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gopherland/db/internal/controller"
	"github.com/gopherland/db/internal/model"
	"github.com/gopherland/db/internal/pg"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	svcPort        = ":3000"
	defaultTimeout = 10 * time.Second
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	db := mustInitDB(ctx)
	defer terminate(db)
	trap(db)

	ma, mb, err := migrate(ctx, db)
	if err != nil {
		log.Panic().Err(err).Msg("migration failed")
	}

	authorH, bookH := controller.NewAuthor(ma), controller.NewBook(mb)

	r := mux.NewRouter()
	pr := r.PathPrefix("/api/v1").Subrouter()
	{
		br := pr.PathPrefix("/books").Subrouter()
		br.HandleFunc("", bookH.Index).Methods("GET")
		br.HandleFunc("/", bookH.Index).Methods("GET")
		br.HandleFunc("/{author}", bookH.ByAuthor).Methods("GET")
	}
	{
		ar := pr.PathPrefix("/authors").Subrouter()
		ar.HandleFunc("", authorH.Index).Methods("GET")
		ar.HandleFunc("/", authorH.Index).Methods("GET")
	}
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(timerWare)

	svc := &http.Server{
		Handler:      r,
		Addr:         svcPort,
		WriteTimeout: defaultTimeout,
		ReadTimeout:  defaultTimeout,
	}
	log.Info().Msgf("BookSvc listening on port %s", svcPort)
	log.Panic().Err(svc.ListenAndServe()).Msgf("service failed")
}

func timerWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func(t time.Time) {
			log.Info().Msgf("[%s] %v %s", r.Method, r.RequestURI, time.Since(t))
		}(time.Now())

		next.ServeHTTP(w, r)
	})
}

func migrate(ctx context.Context, db *sql.DB) (*model.Authors, *model.Books, error) {
	_, err := db.ExecContext(ctx, "drop table if exists books_authors")
	if err != nil {
		return nil, nil, err
	}
	author := model.NewAuthors(db)
	if err := author.Migrate(ctx); err != nil {
		return nil, nil, err
	}
	book := model.NewBooks(db)
	if err := book.Migrate(ctx); err != nil {
		return nil, nil, err
	}
	ba := model.NewBooksAuthors(db)
	if err := ba.Migrate(ctx); err != nil {
		return nil, nil, err
	}

	return author, book, nil
}

func mustInitDB(ctx context.Context) *sql.DB {
	opts := pg.DialOpts{
		<<!!YOUR_CODE!!>> -- define connection options
	}
	db, err := pg.Dial(opts)
	if err != nil {
		panic(err)
	}
	if err = db.PingContext(ctx); err != nil {
		panic(err)
	}

	return db
}

func trap(db *sql.DB) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	go func() {
		<-sig
		log.Debug().Msgf("Terminating!")
		terminate(db)
	}()
}

func terminate(db *sql.DB) {
	if db == nil {
		os.Exit(1)
	}
	if err := db.Close(); err != nil {
		panic(err)
	}
	os.Exit(0)
}

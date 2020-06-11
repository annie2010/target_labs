package main

import (
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gopherland/target_labs/gorm/internal/controller"
	"github.com/gopherland/target_labs/gorm/internal/model"
	"github.com/gopherland/target_labs/gorm/internal/pg"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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
	db := mustInitDB()
	db.LogMode(true)
	defer terminate(db)
	trap(db)

	ma, mb, err := migrate(db)
	if err != nil {
		log.Panic().Err(err)
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

func migrate(db *gorm.DB) (*model.Authors, *model.Books, error) {
	db.DropTableIfExists(&model.BookAuthor{})
	var author model.Author
	db.DropTableIfExists(&author)
	db.AutoMigrate(&author)
	db.Exec("alter sequence authors_id_seq restart with 1 increment by 1")

	var book model.Book
	db.DropTableIfExists(&book)
	db.AutoMigrate(&book)
	db.Exec("alter sequence books_id_seq restart with 1 increment by 1")

	var bookauthor model.BookAuthor
	db.DropTableIfExists(&bookauthor)
	db.AutoMigrate(&bookauthor)

	authors, books := model.NewAuthors(db), model.NewBooks(db)
	if err := authors.Seed(); err != nil {
		return nil, nil, err
	}
	if err := books.Seed(); err != nil {
		return nil, nil, err
	}
	ba := model.NewBooksAuthors(db)
	if err := ba.Seed(); err != nil {
		return nil, nil, err
	}

	return authors, books, nil
}

func mustInitDB() *gorm.DB {
	opts := pg.DialOpts{
		User:     os.Getenv("PG_USER"),
		Password: os.Getenv("PG_PWD"),
		Host:     os.Getenv("PG_HOST"),
		Port:     os.Getenv("PG_PORT"),
		DbName:   os.Getenv("PG_DB"),
	}

	db, err := pg.Dial(opts)
	if err != nil {
		panic(err)
	}

	return db
}

func trap(db *gorm.DB) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	go func() {
		<-sig
		log.Debug().Msgf("Terminating!")
		terminate(db)
	}()
}

func terminate(db *gorm.DB) {
	if db == nil {
		os.Exit(1)
	}
	if err := db.Close(); err != nil {
		panic(err)
	}
	os.Exit(0)
}

func timerWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func(t time.Time) {
			log.Info().Msgf("[%s] %v %s", r.Method, r.RequestURI, time.Since(t))
		}(time.Now())

		next.ServeHTTP(w, r)
	})
}

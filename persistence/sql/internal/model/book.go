// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0
package model

import (
	"context"
	"crypto/sha1"
	"database/sql"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
)

// Select books by author last name query.
const byAuthor = `select * from books b
	where b.id in (
		select book_id from books_authors where author_id in (
			select id from authors a where a.last_name=$1
		)
	);
`

// Book represents a book.
type Book struct {
	ID          int       `json:"id"`
	ISBN        string    `json:"isbn"`
	Title       string    `json:"title"`
	PublishedOn time.Time `json:"published_on"`
}

// Books represents a persistent books model.
type Books struct {
	db           *sql.DB
}

// NewBooks returns a new instance.
func NewBooks(db *sql.DB) *Books {
	return &Books{db: db}
}

// ByAuthor finds all books by a given author last name.
func (b *Books) ByAuthor(ctx context.Context, last string) ([]Book, error) {
	<<!!YOUR_CODE!!>> -- retrieve all books by the given author (hint: checkout byAuthor const above )
}

// Index retrieves all books.
func (b *Books) List(ctx context.Context) ([]Book, error) {
	<<!!YOUR_CODE!!>> -- retrieve books from database
}

const (
	booksDropDDL   = `drop table if exists books;`
	booksCreateDDL = `create table books(
		id serial primary key,
		ISBN varchar(50) unique not null,
		title varchar(100) not null,
		published_on timestamp not null
	);`
	booksInsertDDL = `insert into books (ISBN, title, published_on) values ($1, $2, $3);`
)

// Migrate migrates the database.
func (b *Books) Migrate(ctx context.Context) (err error) {
	log.Debug().Msgf("Migrating Book...")
	txn, err := b.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}
	defer func() {
		if err == nil {
			if err = txn.Commit(); err != nil {
				log.Error().Err(err).Msg("books commit failed")
			}
			return
		}
		log.Error().Err(err).Msg("books migration failed")
		err = txn.Rollback()
	}()

	if _, err = b.db.ExecContext(ctx, booksDropDDL); err != nil {
		return
	}
	if _, err = b.db.ExecContext(ctx, booksCreateDDL); err != nil {
		return
	}

	insertStmt, err := b.db.PrepareContext(ctx, booksInsertDDL)
	if err != nil {
		return err
	}
	defer func() {
		if err = insertStmt.Close(); err != nil {
			return
		}
	}()

	for i := 0; i < 10; i++ {
		title := fmt.Sprintf("Rango%d", i)
		if _, err = insertStmt.ExecContext(
			ctx,
			fmt.Sprintf("%x", sha1.Sum([]byte(title))),
			title,
			time.Now(),
		); err != nil {
			return
		}
	}

	return
}

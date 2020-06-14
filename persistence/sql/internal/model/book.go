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

	gen "github.com/Pallinder/go-randomdata"
	"github.com/rs/zerolog/log"
)

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
	byAuthorStmt *sql.Stmt
}

// NewBooks returns a new instance.
func NewBooks(db *sql.DB) *Books {
	return &Books{db: db}
}

func (b *Books) init(ctx context.Context) (err error) {
  <<!!YOUR_CODE!!>> -- BONUS use a prepared statement to retrive books by authors.
}

func (b *Books) ByAuthor(ctx context.Context, last string) ([]Book, error) {
	<<!!YOUR_CODE!!>> -- Fetch all books by the given author last name (HINT: use the const above or prepared statement)
}

// Index retrieves all books.
func (b *Books) List(ctx context.Context) ([]Book, error) {
	<<!!YOUR_CODE!!>> -- list out all books by querying the db
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
	var txn *sql.Tx
	if txn, err = b.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable}); err != nil {
		return
	}
	defer func() {
		if err != nil {
			err = txn.Rollback()
			return
		}
		log.Info().Msgf("✅ Migrating Book...")
		err = txn.Commit()
	}()

	if _, err = txn.ExecContext(ctx, booksDropDDL); err != nil {
		return
	}
	if _, err = txn.ExecContext(ctx, booksCreateDDL); err != nil {
		return
	}

	insertStmt, err := txn.PrepareContext(ctx, booksInsertDDL)
	if err != nil {
		return err
	}
	defer func() {
		err = insertStmt.Close()
	}()
	for i := 0; i < 10; i++ {
		title := gen.SillyName()
		_, err = insertStmt.ExecContext(ctx, fmt.Sprintf("%x", sha1.Sum([]byte(title))), title, time.Now())
		if err != nil {
			return
		}
	}

	return
}

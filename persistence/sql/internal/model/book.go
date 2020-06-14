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
	if b.byAuthorStmt != nil {
		return err
	}

	const byAuthor = `select * from books b
	  where b.id in (
			select book_id from books_authors where author_id in (
				select id from authors a where a.last_name=$1
			)
		);
	`
	b.byAuthorStmt, err = b.db.PrepareContext(ctx, byAuthor)
	return
}

func (b *Books) ByAuthor(ctx context.Context, last string) ([]Book, error) {
	if err := b.init(ctx); err != nil {
		return nil, err
	}
	rows, err := b.byAuthorStmt.QueryContext(ctx, last)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Error().Err(err).Msgf("closing author rows")
		}
	}()

	bb := make([]Book, 0, 10)
	for rows.Next() {
		var b Book
		if err = rows.Scan(&b.ID, &b.ISBN, &b.Title, &b.PublishedOn); err != nil {
			return nil, err
		}
		bb = append(bb, b)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return bb, nil
}

// Index retrieves all books.
func (b *Books) List(ctx context.Context) ([]Book, error) {
	rows, err := b.db.QueryContext(ctx, "select * from books")
	if err != nil || rows.Err() != nil {
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Error().Err(err).Msgf("closing book rows")
		}
	}()

	bb := make([]Book, 10)
	for rows.Next() {
		var b Book
		if err = rows.Scan(&b.ID, &b.ISBN, &b.Title, &b.PublishedOn); err != nil {
			return nil, err
		}
		bb = append(bb, b)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return bb, nil
}

const (
	booksDropDDL   = `drop table if exists books;`
	booksCreateDDL = `create table books(
		id serial primary key,
		ISBN varchar(50) unique not null,
		title varchar(100) not null,
		published_on timestamp not null
	);`
	booksIndexDDL  = `create index title_idx on books(title);`
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
	if _, err = b.db.ExecContext(ctx, booksIndexDDL); err != nil {
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
		title := gen.SillyName()
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

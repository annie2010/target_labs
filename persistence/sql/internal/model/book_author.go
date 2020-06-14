// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0
package model

import (
	"context"
	"database/sql"

	"github.com/rs/zerolog/log"
)

// BookAuthor represents a join table between books and authors.
type BookAuthor struct {
	BookID   int `json:"book_id"`
	AuthorID int `json:"author_id"`
}

// BooksAuthors represents a book-author model.
type BooksAuthors struct {
	db *sql.DB
}

func NewBooksAuthors(db *sql.DB) BooksAuthors {
	return BooksAuthors{db: db}
}

const (
	baDeleteDDL = `drop table if exists books_authors;`
	baCreateDDL = `create table books_authors(
		book_id int references books(id) on delete cascade,
		author_id int references authors(id) on delete cascade
	);`
	baInsertDDL = `insert into books_authors (book_id, author_id) values ($1, $2);`
)

// Migrate migrates the database.
func (ba BooksAuthors) Migrate(ctx context.Context) error {
	log.Debug().Msgf("Migrating BookAuthor...")
	txn, err := ba.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}
	defer func() {
		if err == nil {
			if err = txn.Commit(); err != nil {
				log.Error().Err(err).Msg("books_authors commit failed")
			}
			return
		}
		log.Error().Err(err).Msg("books_authors migration failed")
		err = txn.Rollback()
	}()

	if _, err := ba.db.ExecContext(ctx, baDeleteDDL); err != nil {
		return err
	}
	if _, err := ba.db.ExecContext(ctx, baCreateDDL); err != nil {
		return err
	}

	insertStmt, err := ba.db.PrepareContext(ctx, baInsertDDL)
	if err != nil {
		return err
	}
	for i := 1; i < 3; i++ {
		if _, err = insertStmt.ExecContext(ctx, i, i+1); err != nil {
			return err
		}
	}
	defer func() {
		if err = insertStmt.Close(); err != nil {
			return
		}
	}()

	return nil
}

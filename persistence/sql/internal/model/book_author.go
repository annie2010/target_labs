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
func (ba BooksAuthors) Migrate(ctx context.Context) (err error) {
	var txn *sql.Tx
	if txn, err = ba.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable}); err != nil {
		return
	}
	defer func() {
		if err != nil {
			err = txn.Rollback()
			return
		}
		log.Info().Msgf("✅ Migrating BooksAuthors...")
		err = txn.Commit()
	}()

	if _, err = txn.ExecContext(ctx, baDeleteDDL); err != nil {
		return
	}
	if _, err = txn.ExecContext(ctx, baCreateDDL); err != nil {
		return
	}

	insertStmt, err := txn.PrepareContext(ctx, baInsertDDL)
	if err != nil {
		return err
	}
	defer func() {
		err = insertStmt.Close()
	}()
	for i := 1; i <= 10; i++ {
		if _, err = insertStmt.ExecContext(ctx, i, i); err != nil {
			return
		}
	}

	return
}

// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0
package model

import (
	"context"
	"database/sql"
	"fmt"
	"math/rand"

	"github.com/rs/zerolog/log"
)

// Author represent an author.
type Author struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

// Authors represents authors model.
type Authors struct {
	db *sql.DB
}

// NewAuthors returns a new instance.
func NewAuthors(db *sql.DB) *Authors {
	return &Authors{db: db}
}

// Index retrieves all Authors.
func (a *Authors) List(ctx context.Context) ([]Author, error) {
	<<!!YOUR_CODE!!>> -- retrieve authors from the database.
}

const (
	authorsDropDDL   = `drop table if exists authors;`
	authorsCreateDDL = `create table authors(
		id serial primary key,
		first_name varchar(30) not null,
		last_name varchar(30) not null,
		age int not null
	);`
	authorsInsertDDL = `insert into authors (first_name, last_name, age) values ($1, $2, $3);`
)

// Migrate migrates the database.
func (a *Authors) Migrate(ctx context.Context) (err error) {
	log.Debug().Msgf("Migrating Authors...")
	txn, err := a.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}
	defer func() {
		if err == nil {
			if err = txn.Commit(); err != nil {
				log.Error().Err(err).Msg("author commit failed")
			}
			return
		}
		log.Error().Err(err).Msg("author migration failed")
		err = txn.Rollback()
	}()

	if _, err = a.db.ExecContext(ctx, authorsDropDDL); err != nil {
		return
	}
	if _, err = a.db.ExecContext(ctx, authorsCreateDDL); err != nil {
		return
	}

	insertStmt, err := a.db.PrepareContext(ctx, authorsInsertDDL)
	if err != nil {
		return err
	}
	defer func() {
		if err = insertStmt.Close(); err != nil {
			return
		}
	}()

	for i := 0; i < 10; i++ {
		if _, err = insertStmt.ExecContext(
			ctx,
			"Fernand",
			fmt.Sprintf("Galiana%d", i),
			20+rand.Int31n(80),
		); err != nil {
			return
		}
	}

	return
}

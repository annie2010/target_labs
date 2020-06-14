// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0
package model

import (
	"context"
	"database/sql"
	"math/rand"

	gen "github.com/Pallinder/go-randomdata"
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
	rows, err := a.db.QueryContext(ctx, "select * from authors")
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Error().Err(err).Msgf("closing book rows")
		}
	}()

	aa := make([]Author, 0, 10)
	for rows.Next() {
		var a Author
		if err = rows.Scan(&a.ID, &a.FirstName, &a.LastName, &a.Age); err != nil {
			return nil, err
		}
		aa = append(aa, a)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return aa, nil
}

const (
	authorsDropDDL   = `drop table if exists authors;`
	authorsCreateDDL = `create table authors(
		id serial primary key,
		first_name varchar(30) not null,
		last_name varchar(30) not null,
		age int not null
	);`
	authorsIndexDDL  = `create index last_idx on authors(last_name);`
	authorsInsertDDL = `insert into authors (first_name, last_name, age) values ($1, $2, $3);`
)

// Migrate migrates the database.
func (a *Authors) Migrate(ctx context.Context) (err error) {
	var txn *sql.Tx
	if txn, err = a.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable}); err != nil {
		return
	}
	defer func() {
		if err != nil {
			err = txn.Rollback()
		}
		log.Info().Msgf("✅ Migrating Authors...")
		err = txn.Commit()
	}()

	if _, err = txn.ExecContext(ctx, authorsDropDDL); err != nil {
		return
	}
	if _, err = txn.ExecContext(ctx, authorsCreateDDL); err != nil {
		return
	}
	if _, err = txn.ExecContext(ctx, authorsIndexDDL); err != nil {
		return
	}

	insertStmt, err := txn.PrepareContext(ctx, authorsInsertDDL)
	if err != nil {
		return err
	}
	defer func() {
		err = insertStmt.Close()
	}()
	for i := 0; i < 10; i++ {
		if _, err = insertStmt.ExecContext(
			ctx,
			gen.FirstName(gen.RandomGender),
			gen.LastName(),
			20+rand.Int31n(80),
		); err != nil {
			return
		}
	}

	return
}

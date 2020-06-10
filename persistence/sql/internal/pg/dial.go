// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0
package pg

import (
	"database/sql"
	"fmt"

	"github.com/rs/zerolog/log"
)

const (
	pgDriver      = "postgres"
	dataSourceFmt = "user=%s password=%s host=%s port=%s dbname=%s sslmode=disable"
)

type DialOpts struct {
	User, Password string
	Host, Port     string
	DbName         string
}

func Dial(opts DialOpts) (*sql.DB, error) {
	ds := fmt.Sprintf(dataSourceFmt,
		opts.User,
		opts.Password,
		opts.Host,
		opts.Port,
		opts.DbName,
	)
	log.Debug().Msgf("DB Connected %s:%s db:%s", opts.Host, opts.Port, opts.DbName)

	return sql.Open(pgDriver, ds)
}

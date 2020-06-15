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

// DialOpts tracks connection configuration.
type DialOpts struct {
	User, Password string
	Host, Port     string
	DbName         string
}

// String returns connection info
func (d DialOpts) String() string {
	return fmt.Sprintf("%s:%s db:%s", d.Host, d.Port, d.DbName)
}

// Flatten returns a datasource string.
func (d DialOpts) flatten() string {
	return fmt.Sprintf(dataSourceFmt,
		d.User,
		d.Password,
		d.Host,
		d.Port,
		d.DbName,
	)
}

// Dial dials in the connection.
func Dial(opts DialOpts) (*sql.DB, error) {
	log.Debug().Msgf("🌏 Connecting DB ... %v", opts)
	return sql.Open(pgDriver, opts.flatten())
}

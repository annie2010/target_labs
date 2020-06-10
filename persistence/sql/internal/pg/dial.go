// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0
package pg

import (
	"database/sql"
	"fmt"
)

const (
	pgDriver      = "postgres"
	dataSourceFmt = "user=%s password=%s host=%s port=%s dbname=%s sslmode=disable"
)

// DialOpts tracks db connection options.
type DialOpts struct {
	User, Password string
	Host, Port     string
	DbName         string
}

// Flatten flattens to data source
func (d DialOpts) Flatten() string {
	<<!!YOUR_CODE!!>> -- return a postgres connection string
}

// Dial configures the DB connection.
func Dial(opts DialOpts) (*sql.DB, error) {
	return sql.Open(pgDriver, opts.Flatten())
}

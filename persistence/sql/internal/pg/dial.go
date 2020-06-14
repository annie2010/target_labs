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
func (d DialOpts) Flatten() string {
	<<!!YOUR_CODE!!>> -- compute a datasource string for the options (HINT: use the const defined above)
}

// Dial dials in the connection.
func Dial(opts DialOpts) (*sql.DB, error) {
	log.Debug().Msgf("🌏 Connecting DB ... %v", opts)
	<<!!YOUR_CODE!!!>> -- open the connection using the correct driver and datasource and return it
}

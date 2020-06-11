package pg

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog/log"
)

const (
	pgDriver      = "postgres"
	dataSourceFmt = "user=%s password=%s host=%s port=%s dbname=%s sslmode=disable"
)

// DialOpts represents db connection config.
type DialOpts struct {
	User, Password string
	Host, Port     string
	DbName         string
}

// String dumps as string.
func (d DialOpts) String() string {
	return fmt.Sprintf("%s:%s db:%s", d.Host, d.Port, d.DbName)
}

func (d DialOpts) flatten() string {
	return fmt.Sprintf(dataSourceFmt,
		d.User,
		d.Password,
		d.Host,
		d.Port,
		d.DbName,
	)
}

// Dial dials the connection.
func Dial(opts DialOpts) (*gorm.DB, error) {
	<<!!YOUR_CODE!!>> -- open a gorm connection to your database
}

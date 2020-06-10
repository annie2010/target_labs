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

type DialOpts struct {
	User, Password string
	Host, Port     string
	DbName         string
}

func Dial(opts DialOpts) (*gorm.DB, error) {
	ds := fmt.Sprintf(dataSourceFmt,
		opts.User,
		opts.Password,
		opts.Host,
		opts.Port,
		opts.DbName,
	)
	log.Debug().Msgf("DB Connected %s:%s db:%s", opts.Host, opts.Port, opts.DbName)

	return gorm.Open(pgDriver, ds)
}

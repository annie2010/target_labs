package model

import (
	"context"
	"math/rand"

	gen "github.com/Pallinder/go-randomdata"
	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog/log"
)

// Author represents an author.
type Author struct {
	gorm.Model

	FirstName string `json:"first_name" gorm:"not null"`
	LastName  string `json:"last_name" gorm:"unique_index;not null"`
	Age       int    `json:"age" gorm:"not null"`
	Books     []Book `gorm:"many2many:books_authors"`
}

// Authors represents authors model.
type Authors struct {
	db *gorm.DB
}

// NewAuthors returns a new instance.
func NewAuthors(db *gorm.DB) *Authors {
	return &Authors{db: db}
}

// Index retrieves all Authors.
func (a *Authors) List(context.Context) ([]Author, error) {
	<<!!YOUR_CODE!!>> - returns all authors from db
}

// Seed seeds the table with fake recs.
func (a *Authors) Seed() error {
	for i := 1; i <= 10; i++ {
		a.db.Create(&Author{
			FirstName: gen.FirstName(gen.RandomGender),
			LastName:  gen.LastName(),
			Age:       int(20 + rand.Int31n(80)),
		})
		if a.db.Error != nil {
			return a.db.Error
		}
	}
	log.Info().Msgf("✅ Migrating Authors...")

	return nil
}

package model

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/jinzhu/gorm"
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
	var aa []Author
	a.db.Find(&aa)

	return aa, a.db.Error
}

// Seed seeds the table with fake recs.
func (a *Authors) Seed() error {
	for i := 1; i <= 10; i++ {
		a.db.Create(&Author{
			FirstName: "Fernand",
			LastName:  fmt.Sprintf("Galiana%d", i),
			Age:       int(20 + rand.Int31n(80)),
		})
		if a.db.Error != nil {
			return a.db.Error
		}
	}
	return nil
}

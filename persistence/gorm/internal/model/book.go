package model

import (
	"context"
	"crypto/sha1"
	"fmt"
	"time"

	gen "github.com/Pallinder/go-randomdata"
	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog/log"
)

// Book represents a book.
type Book struct {
	gorm.Model

	ISBN        string    `json:"isbn" gorm:"not null"`
	Title       string    `json:"title" gorm:"index, not null"`
	PublishedOn time.Time `json:"published_on"`
	Authors     []Author  `gorm:"many2many:books_authors"`
}

// Books represents a persistent books model.
type Books struct {
	db *gorm.DB
}

// NewBooks returns a new instance.
func NewBooks(db *gorm.DB) *Books {
	return &Books{db: db}
}

// ByAuthor returns books from a given author.
func (b *Books) ByAuthor(_ context.Context, last string) ([]Book, error) {
	<<!!YOUR_CODE!!>> - returns all books from a given author last name
}

// Index retrieves all books.
func (b *Books) List(context.Context) ([]Book, error) {
	<<!!YOUR_CODE!!>> - return all books from db
}

// Seed seeds the table with fake recs.
func (b *Books) Seed() error {
	for i := 1; i <= 10; i++ {
		title := gen.SillyName()
		b.db.Create(&Book{
			ISBN:        fmt.Sprintf("%x", sha1.Sum([]byte(title))),
			Title:       title,
			PublishedOn: time.Now(),
		})
		if b.db.Error != nil {
			return b.db.Error
		}
	}
	log.Info().Msgf("✅ Migrating Books...")

	return nil
}

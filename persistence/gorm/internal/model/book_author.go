package model

import (
	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog/log"
)

// BookAuthor represents a join table.
type BookAuthor struct {
	BookID   int
	AuthorID int
}

// TableName overrides the table name
func (BookAuthor) TableName() string {
	return "books_authors"
}

// BooksAuthors represents an association.
type BooksAuthors struct {
	db *gorm.DB
}

// NewBooksAuthors returns a new instance.
func NewBooksAuthors(db *gorm.DB) *BooksAuthors {
	return &BooksAuthors{
		db: db,
	}
}

// Seed seeds the table.
func (ba *BooksAuthors) Seed() error {
	for i := 1; i < 3; i++ {
		ba.db.Create(&BookAuthor{
			BookID:   i,
			AuthorID: i + 1,
		})
	}
	log.Info().Msgf("✅ Migrating BooksAuthors...")

	return ba.db.Error
}

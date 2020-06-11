package model

import (
	"context"
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
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

// ByAuthor fetch a books from the given author last name.
func (b *Books) ByAuthor(ctx context.Context, last string) ([]Book, error) {
	<<!!YOUR_CODE!!>> -- fetch all books from the given author
}

// Index retrieves all books.
func (b *Books) List(ctx context.Context) ([]Book, error) {
	<<!!YOUR_CODE!!>> -- fetch all books from DB
}

// Seed seeds the table.
func (b *Books) Seed() error {
	for i := 1; i <= 10; i++ {
		title := fmt.Sprintf("Rango%d", i)
		b.db.Create(&Book{
			ISBN:        fmt.Sprintf("%x", sha1.Sum([]byte(title))),
			Title:       title,
			PublishedOn: time.Now(),
		})
		if b.db.Error != nil {
			return b.db.Error
		}
	}

	return nil
}

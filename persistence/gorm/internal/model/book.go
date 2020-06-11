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

func (b *Books) ByAuthor(_ context.Context, last string) ([]Book, error) {
	var bb []Book
	b.db.Where("id in ?",
		b.db.Table("books_authors").Unscoped().Select("book_id").
			Where("author_id in ?",
				b.db.Table("authors").Select("id").Where("last_name = ?", last).SubQuery(),
			).SubQuery(),
	).Find(&bb)

	return bb, b.db.Error
}

// Index retrieves all books.
func (b *Books) List(context.Context) ([]Book, error) {
	var bb []Book
	b.db.Find(&bb)

	return bb, b.db.Error
}

// Seed seeds the table with fake recs.
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

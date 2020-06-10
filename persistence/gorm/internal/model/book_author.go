package model

import "github.com/jinzhu/gorm"

type BookAuthor struct {
	BookID   int
	AuthorID int
}

func (BookAuthor) TableName() string {
	return "books_authors"
}

type BooksAuthors struct {
	db *gorm.DB
}

func NewBooksAuthors(db *gorm.DB) *BooksAuthors {
	return &BooksAuthors{
		db: db,
	}
}

func (ba *BooksAuthors) Seed() error {
	for i := 1; i < 3; i++ {
		ba.db.Create(&BookAuthor{
			BookID:   i,
			AuthorID: i + 1,
		})
	}
	return ba.db.Error
}

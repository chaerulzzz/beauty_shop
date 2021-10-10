package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title  string
	Author string
}

type BookDTO struct {
	ID     uint   `json:"id,string,omitempty"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func ToBook(dto BookDTO) Book {
	return Book{Title: dto.Title, Author: dto.Author}
}

func ToBookDTO(book Book) BookDTO {
	return BookDTO{ID: book.ID, Title: book.Title, Author: book.Author}
}

func ToBookDTOs(books []Book) []BookDTO {
	booksDTOs := make([]BookDTO, len(books))

	for i, item := range books {
		booksDTOs[i] = ToBookDTO(item)
	}

	return booksDTOs
}

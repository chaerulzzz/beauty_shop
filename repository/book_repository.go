package repository

import (
	"beauty-shop/models"
	"github.com/go-redis/redis/v7"
)

type BookRepository struct {
	repo BaseRepository
}

func ProvideBookRepository(repo BaseRepository) BookRepository {
	return BookRepository{repo: repo}
}

func (p *BookRepository) FindAll() []models.Book {
	var books []models.Book
	p.repo.DB.Find(&books)

	return books
}

func (p *BookRepository) FindByID(id uint) models.Book {
	var book models.Book
	p.repo.DB.Find(&book, id)

	return book
}

func (p *BookRepository) Save(book models.Book) models.Book {
	p.repo.DB.Save(&book)

	return book
}

func (p *BookRepository) Delete(book models.Book) {
	p.repo.DB.Delete(&book)
}

func (p *BookRepository) GetRedisClient() *redis.Client {
	return p.repo.client
}

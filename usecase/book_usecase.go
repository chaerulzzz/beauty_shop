package usecase

import (
	"beauty-shop/models"
	"beauty-shop/repository"
	"github.com/go-redis/redis/v7"
)

type BookUsecase struct {
	BookRepository repository.BookRepository
}

func ProvideBookUseCase(p repository.BookRepository) BookUsecase {
	return BookUsecase{BookRepository: p}
}

func (p *BookUsecase) FindAll() []models.Book {
	return p.BookRepository.FindAll()
}

func (p *BookUsecase) FindByID(id uint) models.Book {
	return p.BookRepository.FindByID(id)
}

func (p *BookUsecase) Save(book models.Book) models.Book {
	p.BookRepository.Save(book)

	return book
}

func (p *BookUsecase) Delete(book models.Book) {
	p.BookRepository.Delete(book)
}

func (p *BookUsecase) GetRedisClient() *redis.Client {
	return p.BookRepository.GetRedisClient()
}

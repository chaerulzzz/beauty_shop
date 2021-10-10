package usecase

import (
	"beauty-shop/models"
	"beauty-shop/repository"
	"github.com/go-redis/redis/v7"
)

type ProductUseCase struct {
	ProductRepository repository.ProductRepository
}

func ProvideProductUseCase(p repository.ProductRepository) ProductUseCase {
	return ProductUseCase{ProductRepository: p}
}

func (p *ProductUseCase) FindAll() []models.Product {
	return p.ProductRepository.FindAllProduct()
}

func (p *ProductUseCase) FindByID(id uint) models.Product {
	return p.ProductRepository.FindByID(id)
}

func (p *ProductUseCase) FindByName(name string) []models.Product {
	return p.ProductRepository.FindByName(name)
}

func (p *ProductUseCase) SaveProduct(product models.Product) models.Product {
	p.ProductRepository.SaveProduct(product)

	return product
}

func (p *ProductUseCase) DeleteProduct(product models.Product) {
	p.ProductRepository.DeleteProduct(product)
}

func (p *ProductUseCase) GetRedisClient() *redis.Client {
	return p.ProductRepository.GetRedisClient()
}

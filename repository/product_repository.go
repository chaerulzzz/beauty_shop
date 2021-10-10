package repository

import (
	"beauty-shop/models"
	"github.com/go-redis/redis/v7"
)

type ProductRepository struct {
	repo BaseRepository
}

func ProvideProductRepository(repo BaseRepository) ProductRepository {
	return ProductRepository{repo: repo}
}

func (p *ProductRepository) FindByID(id uint) models.Product {
	var product models.Product
	p.repo.DB.Find(&product, id)

	return product
}

func (p *ProductRepository) FindByName(name string) []models.Product {
	var products []models.Product
	p.repo.DB.Where("productName ilike %?%", name).Find(&products)

	return products
}

func (p *ProductRepository) FindAllProduct() []models.Product {
	var products []models.Product
	p.repo.DB.Find(&products)

	return products
}

func (p *ProductRepository) SaveProduct(product models.Product) models.Product {
	p.repo.DB.Save(&product)

	return product
}

func (p *ProductRepository) DeleteProduct(product models.Product) {
	p.repo.DB.Delete(&product)
}

func (p *ProductRepository) GetRedisClient() *redis.Client {
	return p.repo.client
}

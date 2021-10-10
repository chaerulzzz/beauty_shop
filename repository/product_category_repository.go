package repository

import "beauty-shop/models"

type ProductCategoryRepository struct {
	repo BaseRepository
}

func ProvideProductCategory(repo BaseRepository) ProductCategoryRepository {
	return ProductCategoryRepository{repo: repo}
}

func (p *ProductCategoryRepository) FindAll() []models.ProductCategory {
	var categories []models.ProductCategory
	p.repo.DB.Find(&categories)

	return categories
}

func (p *ProductCategoryRepository) Save(category models.ProductCategory) models.ProductCategory {
	p.repo.DB.Save(&category)

	return category
}

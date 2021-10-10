package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string
	Description string
	Image       string
	Price       uint32
	CategoryId  uint
}

type ProductDTO struct {
	ID          uint   `json:"id,omitempty"`
	Name        string `json:"productName"`
	Description string `json:"productDescription"`
	Image       string `json:"productImage,omitempty"`
	Price       uint32 `json:"productPrice"`
	CategoryId  uint   `json:"categoryId,omitempty"`
}

func ToProduct(dto ProductDTO) Product {
	return Product{Name: dto.Name, Description: dto.Description, Image: dto.Image, Price: dto.Price, CategoryId: dto.CategoryId}
}

func ToProductDTO(product Product) ProductDTO {
	return ProductDTO{ID: product.ID, Name: product.Name, Description: product.Description, Image: product.Image, Price: product.Price, CategoryId: product.CategoryId}
}

func ToProductDTOs(products []Product) []ProductDTO {
	booksDTOs := make([]ProductDTO, len(products))

	for i, item := range products {
		booksDTOs[i] = ToProductDTO(item)
	}

	return booksDTOs
}

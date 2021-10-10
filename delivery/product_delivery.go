package delivery

import (
	"beauty-shop/auth"
	"beauty-shop/models"
	"beauty-shop/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	ProductCase usecase.ProductUseCase
}

func ProvideProductHandler(p usecase.ProductUseCase) ProductHandler {
	return ProductHandler{ProductCase: p}
}

func (p *ProductHandler) GetAll(c *gin.Context) {
	tokenAuth, err := auth.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	_, err = auth.FetchAuth(tokenAuth, p.ProductCase.GetRedisClient())
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	products := p.ProductCase.FindAll()

	if len(products) <= 0 {
		c.JSON(http.StatusOK, gin.H{
			"responseCode":    1,
			"responseMessage": "Produk tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"responseCode":    0,
		"responseMessage": "Sukses",
		"products":        models.ToProductDTOs(products)})
}

func (p *ProductHandler) GetProductByName(c *gin.Context) {
	tokenAuth, err := auth.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	_, err = auth.FetchAuth(tokenAuth, p.ProductCase.GetRedisClient())
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	var u models.ProductDTO
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"responseMessage": err.Error(),
		})
		return
	}

	products := p.ProductCase.FindByName(u.Name)
	if len(products) <= 0 {
		c.JSON(http.StatusOK, gin.H{
			"responseCode":    1,
			"responseMessage": "Produk tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"responseCode":    0,
		"responseMessage": "sukses",
		"products":        models.ToProductDTOs(products),
	})
}

func (p *ProductHandler) GetById(c *gin.Context) {
	tokenAuth, err := auth.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	_, err = auth.FetchAuth(tokenAuth, p.ProductCase.GetRedisClient())
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Product ID"})
		return
	}

	product := p.ProductCase.FindByID(uint(id))
	if product == (models.Product{}) {
		c.JSON(http.StatusOK, gin.H{
			"responseCode":    1,
			"responseMessage": "Produk tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"responseCode":    0,
		"responseMessage": "Produk tidak ditemukan",
		"product":         models.ToProductDTO(product),
	})
}

func (p *ProductHandler) Create(c *gin.Context) {
	tokenAuth, errs := auth.ExtractTokenMetadata(c.Request)
	if errs != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	_, errs = auth.FetchAuth(tokenAuth, p.ProductCase.GetRedisClient())
	if errs != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	var productDto models.ProductDTO
	err := c.BindJSON(&productDto)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	_ = p.ProductCase.SaveProduct(models.ToProduct(productDto))

	c.JSON(http.StatusOK, gin.H{
		"responseCode":    0,
		"responseMessage": "Produk sukses dibuat"})
}

func (p *ProductHandler) Update(c *gin.Context) {
	tokenAuth, errs := auth.ExtractTokenMetadata(c.Request)
	if errs != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	_, errs = auth.FetchAuth(tokenAuth, p.ProductCase.GetRedisClient())
	if errs != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	var productDto models.ProductDTO
	err := c.BindJSON(&productDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"responseCode":    2,
			"responseMessage": "ID product invalid"})
		return
	}

	product := p.ProductCase.FindByID(uint(id))
	if product == (models.Product{}) {
		c.JSON(http.StatusOK, gin.H{
			"responseCode":    1,
			"responseMessage": "Produk tidak ditemukan"})
		return
	}

	product.Name = productDto.Name
	product.Description = productDto.Description
	product.Price = productDto.Price
	product.Image = productDto.Image
	p.ProductCase.SaveProduct(product)

	c.JSON(http.StatusOK, gin.H{
		"responseCode":    0,
		"responseMessage": "Sukses menyimpan produk"})
}

func (p *ProductHandler) Delete(c *gin.Context) {
	tokenAuth, err := auth.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	_, err = auth.FetchAuth(tokenAuth, p.ProductCase.GetRedisClient())
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Book ID"})
		return
	}

	product := p.ProductCase.FindByID(uint(id))
	if product == (models.Product{}) {
		c.JSON(http.StatusOK, gin.H{
			"responseCode":    1,
			"responseMessage": "Produk tidak ditemukan",
		})
		return
	}

	p.ProductCase.DeleteProduct(product)
	c.JSON(http.StatusOK, gin.H{
		"responseCode":    0,
		"responseMessage": "Sukses menghapus produk"})
}

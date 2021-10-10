package delivery

import (
	"beauty-shop/auth"
	"beauty-shop/models"
	"beauty-shop/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type BookHandler struct {
	BookService usecase.BookUsecase
}

func ProviderBookHandler(p usecase.BookUsecase) BookHandler {
	return BookHandler{BookService: p}
}

func (p *BookHandler) FindAll(c *gin.Context) {
	tokenAuth, err := auth.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	_, err = auth.FetchAuth(tokenAuth, p.BookService.GetRedisClient())
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	books := p.BookService.FindAll()

	if len(books) <= 0 {
		c.JSON(http.StatusOK, gin.H{
			"books":   models.ToBookDTOs(books),
			"message": "List is empty!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"books": models.ToBookDTOs(books)})
}

func (p *BookHandler) FindByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Book ID"})
		return
	}

	book := p.BookService.FindByID(uint(id))
	if book == (models.Book{}) {
		c.JSON(http.StatusOK, gin.H{
			"book":    "{}",
			"message": "Book not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"book": models.ToBookDTO(book)})
}

func (p *BookHandler) Create(c *gin.Context) {
	var bookDTO models.BookDTO
	err := c.BindJSON(&bookDTO)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	createBook := p.BookService.Save(models.ToBook(bookDTO))

	c.JSON(http.StatusCreated, gin.H{"book": models.ToBookDTO(createBook)})
}

func (p *BookHandler) Update(c *gin.Context) {
	var bookDTO models.BookDTO
	err := c.BindJSON(&bookDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Book ID"})
		return
	}

	book := p.BookService.FindByID(uint(id))
	if book == (models.Book{}) {
		c.JSON(http.StatusOK, gin.H{"message": "Book not found!"})
		return
	}

	book.Author = bookDTO.Author
	book.Title = bookDTO.Title
	p.BookService.Save(book)

	c.JSON(http.StatusNoContent, gin.H{"message": ""})
}

func (p *BookHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Book ID"})
		return
	}

	book := p.BookService.FindByID(uint(id))
	if book == (models.Book{}) {
		c.JSON(http.StatusOK, gin.H{
			"book":    "{}",
			"message": "Book not found",
		})
		return
	}

	p.BookService.Delete(book)
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}

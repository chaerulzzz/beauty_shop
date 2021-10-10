package delivery

import (
	"beauty-shop/auth"
	"beauty-shop/helper"
	"beauty-shop/models"
	"beauty-shop/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	UserCase usecase.UserUsecase
}

func ProviderUserHandler(p usecase.UserUsecase) UserHandler {
	return UserHandler{UserCase: p}
}

func (p *UserHandler) GetUserByEmail(c *gin.Context) {
	var u models.UserDTO
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	user := p.UserCase.FindByEmail(u.Email)
	if user == (models.User{}) {
		c.JSON(http.StatusNotFound, gin.H{
			"user":    "{}",
			"message": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": models.ToUserDTO(user)})
}

func (p *UserHandler) GetUserByUsername(c *gin.Context) {
	var u models.UserDTO
	if err := c.ShouldBindJSON(&u); err != nil {
		helper.ResponseJSON(c, http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	user := p.UserCase.FindByUsername(u.Username)
	if user == (models.User{}) {
		c.JSON(http.StatusNotFound, gin.H{
			"user":    "{}",
			"message": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": models.ToUserDTO(user)})
}

func (p *UserHandler) Register(c *gin.Context) {
	var userDTO models.UserDTO
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	if err := helper.ValidateEmail(userDTO.Email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if len(userDTO.Email) <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Email tidak boleh kosong"})
		return
	}

	if len(userDTO.Username) <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Username tidak boleh kosong"})
		return
	}

	if len(userDTO.Password) <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Password tidak boleh kosong"})
		return
	}

	registerUser := p.UserCase.SaveUser(models.ToUser(userDTO))

	c.JSON(http.StatusOK, gin.H{"user": models.ToUserDTO(registerUser)})
}

func (p *UserHandler) Login(c *gin.Context) {
	var userDTO models.UserDTO
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid")
	}

	user := p.UserCase.FindByUsernameAndPassword(userDTO.Username, userDTO.Password)
	if user == (models.User{}) {
		c.JSON(http.StatusOK, gin.H{
			"responseCode":    1,
			"responseMessage": "Username atau password salah",
		})
		return
	}

	token, err := auth.CreateToken(user.ID)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	saveErr := auth.CreateAuth(user.ID, token, p.UserCase.UserRepository.GetRedisClient())
	if saveErr != nil {
		c.JSON(http.StatusUnprocessableEntity, saveErr.Error())
	}

	var roles string
	if user.Username == "admin" {
		roles = "administrator"
	} else {
		roles = "customer"
	}

	tokens := map[string]interface{}{
		"responseCode":    0,
		"responseMessage": "Sukses",
		"roles":           roles,
		"accessToken":     token.AccessToken,
		"refreshToken":    token.RefreshToken,
	}

	c.JSON(http.StatusOK, tokens)
}

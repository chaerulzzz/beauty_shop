package repository

import (
	"beauty-shop/models"
	"github.com/go-redis/redis/v7"
)

type UserRepository struct {
	repo BaseRepository
}

func ProvideUserRepository(repo BaseRepository) UserRepository {
	return UserRepository{repo: repo}
}

func (p *UserRepository) FindByID(id uint) models.User {
	var user models.User
	p.repo.DB.Find(&user, id)

	return user
}

func (p *UserRepository) FindByByUsername(username string) models.User {
	var user models.User
	p.repo.DB.Where("username = ?", username).First(&user)

	return user
}

func (p *UserRepository) FindByByUsernameAndPassword(username string, password string) models.User {
	var user models.User
	p.repo.DB.Where("username = ? AND password = ?", username, password).First(&user)

	return user
}

func (p *UserRepository) FindByEmail(email string) models.User {
	var user models.User
	p.repo.DB.Where("email = ?", email).First(&user)

	return user
}

func (p *UserRepository) SaveUser(user models.User) models.User {
	p.repo.DB.Save(&user)

	return user
}

func (p *UserRepository) GetRedisClient() *redis.Client {
	return p.repo.client
}

package usecase

import (
	"beauty-shop/models"
	"beauty-shop/repository"
)

type UserUsecase struct {
	UserRepository repository.UserRepository
}

func ProvideUserUseCase(p repository.UserRepository) UserUsecase {
	return UserUsecase{UserRepository: p}
}

func (p *UserUsecase) FindById(id uint) models.User {
	return p.UserRepository.FindByID(id)
}

func (p *UserUsecase) FindByUsername(username string) models.User {
	return p.UserRepository.FindByByUsername(username)
}

func (p *UserUsecase) FindByUsernameAndPassword(username string, password string) models.User {
	return p.UserRepository.FindByByUsernameAndPassword(username, password)
}

func (p *UserUsecase) FindByEmail(email string) models.User {
	return p.UserRepository.FindByEmail(email)
}

func (p *UserUsecase) SaveUser(user models.User) models.User {
	p.UserRepository.SaveUser(user)

	return user
}

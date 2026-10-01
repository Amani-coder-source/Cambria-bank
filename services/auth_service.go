package services

import (
	"errors"

	"cambria-bank/models"
	"cambria-bank/repository"
	"cambria-bank/utils"
)

type AuthService struct {
	Repo        repository.UserRepository
	AccountRepo repository.AccountRepository
}

func (a AuthService) Register(user models.User) error {

	if user.Role == "" {
		user.Role = "customer"
	}

	hash, err := utils.HashPassword(user.Password)

	if err != nil {
		return err
	}

	user.Password = hash

	userID, err := a.Repo.CreateUser(user)

	if err != nil {
		return err
	}

	accountService := AccountService{
		Repo: a.AccountRepo,
	}

	err = accountService.CreateAccount(userID)

	if err != nil {
		return err
	}

	return nil
}

func (a AuthService) Login(email string, password string) (string, error) {

	user, err := a.Repo.GetUserByEmail(email)

	if err != nil {
		return "", errors.New("email not found")
	}

	if !utils.CheckPassword(password, user.Password) {
		return "", errors.New("wrong password")
	}

	token, err := utils.GenerateToken(user.ID, user.Email)

	if err != nil {
		return "", err
	}

	return token, nil
}

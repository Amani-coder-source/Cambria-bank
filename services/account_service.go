package services

import (
	"errors"
	"fmt"
	"strconv"

	"cambria-bank/models"
	"cambria-bank/repository"
)

type AccountService struct {
	Repo repository.AccountRepository
}

func (a AccountService) CreateAccount(userID int) error {

	accountNumber := "CB" + fmt.Sprintf("%08d", userID)

	account := models.Account{
		UserID:        userID,
		AccountNumber: accountNumber,
		AccountType:   "current",
		Balance:       2500000.00,
		Currency:      "GBP",
	}

	return a.Repo.CreateAccount(account)
}

func (a AccountService) GetAccount(userID int) (models.Account, error) {

	account, err := a.Repo.GetAccountByUserID(userID)

	if err != nil {
		return models.Account{}, errors.New("account not found")
	}

	return account, nil
}

func GenerateAccountNumber(userID int) string {

	return "CB" + strconv.Itoa(userID)
}

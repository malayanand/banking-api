package domain

import (
	"github.com/malayanand/banking/dto"
	"github.com/malayanand/banking/errs"
)

type Account struct {
	AccountId   string  `db:"account_id"`
	CustomerId  string  `db:"customer_id"`
	OpeningDate string  `db:"opening_date"`
	AccountType string  `db:"account_type"`
	Amount      float64 `db:"amount"`
	Status      string  `db:"status"`
}

func (a Account) ToNewAccountResponseDto() *dto.NewAccountResponse {
	return &dto.NewAccountResponse{AccountId: a.AccountId}
}

func (a Account) CanWithdraw(amount float64) bool {
	if a.Amount >= amount {
		return true
	}
	return false
}

type AccountRepository interface {
	FindBy(accountId string) (*Account, *errs.AppError)
	SaveTransaction(transaction Transaciton) (*Transaciton, *errs.AppError)
	Save(Account) (*Account, *errs.AppError)
}

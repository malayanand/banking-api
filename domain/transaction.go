package domain

import "github.com/malayanand/banking/dto"

type Transaciton struct {
	TransactionId   string
	AccountId       string
	TransactionType string
	TransactionDate string
	Amount          float64
}

func (t Transaciton) ToDto() *dto.TransactionResponse {
	return &dto.TransactionResponse{
		TransactionId:   t.TransactionId,
		AccountId:       t.AccountId,
		Amount:          t.Amount,
		TransactionType: t.TransactionType,
		TransactionDate: t.TransactionDate,
	}
}

func (t Transaciton) IsWitdrawal() bool {
	if t.TransactionType == "withdrawal" {
		return true
	}
	return false
}

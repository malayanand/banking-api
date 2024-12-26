package service

import (
	"time"

	"github.com/malayanand/banking/domain"
	"github.com/malayanand/banking/dto"
	"github.com/malayanand/banking/errs"
)

const dbTSLayout = "2006-01-02 15:04:05"

type DefaultAccountService struct {
	repo domain.AccountRepository
}

type AccountService interface {
	MakeTransaction(dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError)
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
}

func (s DefaultAccountService) MakeTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	// check for correct account balance before withdrwawal
	if req.IsTransactionWithdrawal() {
		account, err := s.repo.FindBy(req.AccountId)
		if err != nil {
			return nil, err
		}

		if !account.CanWithdraw(req.Amount) {
			return nil, errs.NewValidationError("Insufficient balance in the account")
		}
	}

	t := domain.Transaciton{
		AccountId:       req.AccountId,
		Amount:          req.Amount,
		TransactionType: req.TransactionType,
		TransactionDate: time.Now().Format(dbTSLayout),
	}

	transaction, appErr := s.repo.SaveTransaction(t)
	if appErr != nil {
		return nil, appErr
	}

	response := transaction.ToDto()
	return response, nil
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	// validate if the request body is correct
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	a := domain.Account{
		AccountId:   "",
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format(time.RFC3339),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}
	newAccount, err := s.repo.Save(a)
	if err != nil {
		return nil, err
	}

	response := newAccount.ToNewAccountResponseDto()

	return response, nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo}
}

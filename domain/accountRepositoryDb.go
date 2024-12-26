package domain

import (
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/malayanand/banking/errs"
	"github.com/malayanand/banking/logger"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {
	sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) values (?,?,?,?,?)"

	result, err := d.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last inserted id for new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	a.AccountId = strconv.FormatInt(id, 10)
	return &a, nil
}

func (d AccountRepositoryDb) FindBy(accountId string) (*Account, *errs.AppError) {
	var account Account
	err := d.client.Get(&account, "SELECT * FROM accounts where account_id = ?", accountId)

	if err != nil {
		logger.Error("Error while fetching the account information " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return &account, nil
}

func (d AccountRepositoryDb) SaveTransaction(t Transaciton) (*Transaciton, *errs.AppError) {
	tx, err := d.client.Begin()
	if err != nil {
		logger.Error("Error while starting a transaction for bank account transfer " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	//inserting bank account transaction
	result, _ := tx.Exec("INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) values (?, ?, ?, ?)", t.AccountId, t.Amount, t.TransactionType, t.TransactionDate)

	if t.IsWitdrawal() {
		_, err = tx.Exec("UPDATE accounts SET amount = amount - ? where account_id = ?", t.Amount, t.AccountId)
	} else {
		_, err = tx.Exec("UPDATE accounts SET amount = amount + ? where account_id = ?", t.Amount, t.AccountId)
	}

	// rollback, in case of error
	if err != nil {
		tx.Rollback()
		logger.Error("Error while saving transaction " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	// commit changes into the database
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Error while commiting the transaction for the account " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	transactionId, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while fetching the last transaction id " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	// fetch the account information
	account, appErr := d.FindBy(t.AccountId)
	if appErr != nil {
		return nil, appErr
	}

	t.TransactionId = strconv.FormatInt(transactionId, 10)
	t.Amount = account.Amount

	return &t, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}

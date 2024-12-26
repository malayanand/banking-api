package domain

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/malayanand/banking/errs"
	"github.com/malayanand/banking/logger"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	findById := "select * from customers where customer_id=?"

	//row := d.client.QueryRow(findById, id)
	var c Customer
	err := d.client.Get(&c, findById, id)

	//err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			logger.Error("Error while scanning customer " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return &c, nil
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	//var rows *sql.Rows
	var err error
	customers := make([]Customer, 0)

	if status == "" {
		findAllSql := "select * from customers"
		err = d.client.Select(&customers, findAllSql)
		//rows, err = d.client.Query(findAllSql)
	} else {
		findAllSql := "select * from customers where status=?"
		err = d.client.Select(&customers, findAllSql, status)
		//rows, err = d.client.Query(findAllSql, status)
	}

	if err != nil {
		logger.Error("Error while querying customer table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	/*
		err = sqlx.StructScan(rows, &customers)
		if err != nil {
			logger.Error("Error while scanning customers " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}*/

	// default boiler plate to store values from query into go struct
	/*
		for rows.Next() {
			var c Customer
			err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
			if err != nil {
				logger.Error("Error while scanning customers " + err.Error())
				return nil, errs.NewUnexpectedError("Unexpected database error")
			}
			customers = append(customers, c)
		}*/

	return customers, nil
}

func NewCustomerRepositoryDb(dbClient *sqlx.DB) CustomerRepositoryDb {
	return CustomerRepositoryDb{dbClient}
}

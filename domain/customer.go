package domain

import (
	"github.com/malayanand/banking/dto"
	"github.com/malayanand/banking/errs"
)

// domain with port defined
type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string `db:"date_of_birth"`
	Status      string
}

func (c Customer) statusAsText() string {
	statusAsText := "active"
	// return active/inactive instead of 0/1
	if c.Status == "0" {
		statusAsText = "inactive"
	}

	return statusAsText
}

func (c Customer) ToDto() *dto.CustomerResponse {
	return &dto.CustomerResponse{
		Id:          c.Id,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateOfBirth: c.DateOfBirth,
		Status:      c.statusAsText(),
	}
}

type CustomerRepository interface {
	// status == 1 status == 0 status == " "
	FindAll(status string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}

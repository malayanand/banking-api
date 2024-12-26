package service

import (
	"github.com/malayanand/banking/domain"
	"github.com/malayanand/banking/dto"
	"github.com/malayanand/banking/errs"
)

// port
type CustomerService interface {
	GetAllCustomers(string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
}

// implementation
type DefaultCustomerService struct {
	repo domain.CustomerRepository // dependency
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]dto.CustomerResponse, *errs.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}

	customers, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err

	}
	response := make([]dto.CustomerResponse, 0)
	for _, c := range customers {
		response = append(response, *c.ToDto())
	}

	return response, nil
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}

	// c.ToDto returns a pointer to DTO
	response := c.ToDto()
	return response, nil
}

// function to intialize customer service
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}

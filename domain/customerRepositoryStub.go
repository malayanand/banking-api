package domain

// adapter
type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "Malay", "New Delhi", "781007", "1998-06-04", "1"},
		{"1002", "Ashish", "New Delhi", "781009", "1998-06-04", "1"},
	}

	return CustomerRepositoryStub{customers}
}

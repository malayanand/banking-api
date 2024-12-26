package dto

type CustomerResponse struct {
	Id          string `json:"customer_id"`
	City        string `json:"city"`
	Name        string `json:"name"`
	Zipcode     string `json:"zip_code"`
	DateOfBirth string `json:"date_of_birth"`
	Status      string `json:"status"`
}

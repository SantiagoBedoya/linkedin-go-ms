package models

type Customer struct {
	CustomerID string `gorm:"primaryKey" json:"customerId"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
}

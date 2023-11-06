package models

type Service struct {
	ServiceID string  `gorm:"primaryKey" json:"serviceId"`
	Name      string  `json:"name"`
	Price     float64 `gorm:"type:numeric" json:"price"`
}

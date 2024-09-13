package models

type ContactInfo struct {
	ID        int    `json:"id" gorm:"primary_key"`
	FirstName string `json:"first_name" gorm:"not null"`
	LastName  string `json:"last_name" gorm:"not null"`
	Email     string `json:"email" gorm:"not null"`
	Phone     string `json:"phone" gorm:"not null"`
}

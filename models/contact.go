package models

type ContactInfo struct {
	ID        int    `json:"id" gorm:"primary_key"`
	FirstName string `json:"name" gorm:"not null"`
	LastName  string `json:"surname" gorm:"not null"`
	Email     string `json:"email" gorm:"unique;not null"`
	Phone     string `json:"phone" gorm:"unique;not null"`
}
package models

import "time"

type ContactInfo struct {
	ID        int       `json:"id" gorm:"primary_key"`
	FirstName string    `json:"first_name" gorm:"not null"`
	LastName  string    `json:"last_name" gorm:"not null"`
	Email     string    `json:"email" gorm:"not null"`
	Phone     string    `json:"phone" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

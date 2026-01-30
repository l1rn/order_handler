package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID		 uuid.UUID	
	Username string
	Password string 
}

func (u *User) BeforeCreate(rx *gorm.DB) (err error) {
	
}
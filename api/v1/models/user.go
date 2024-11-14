package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName  string `gorm:"column:first_name;not null" json:"first_name"`
	MiddleName string `gorm:"column:middle_name" json:"middle_name"`
	LastName   string `gorm:"column:last_name;not null" json:"last_name"`
	Email      string `gorm:"column:email;unique" json:"email"`
	Username   string `gorm:"column:username;unique" json:"username"`
	Password   string `gorm:"column:password;not null" json:"password"`
}

package model

type User struct {
	Base
	Username string `gorm:"unique"`
	Name     string
}

package model

type User struct {
	Id       int
	Name     string
	Username string `gorm:"primaryKey"`
	Password string
}

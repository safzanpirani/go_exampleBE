package dao

import (
	"errors"
	"fmt"
	"gin/basic/model"
	"gorm.io/gorm"
)

func SaveUser(db *gorm.DB, user *model.User) error {
	err := db.Save(user).Error
	return err
}
func LogIn(db *gorm.DB, user *model.User) error {
	//get user by username from db (see below, superseded)
	// db.password==user.password

	getUser := model.User{}
	err := db.Select("users.*").Where("users.username=?", user.Username).First(&getUser).Error
	if err != nil {
		return fmt.Errorf("User not found, getlost, or try again", err.Error())
	}
	if user.Password != getUser.Password {
		return errors.New("Password Incorrect")
	}
	return nil
}

func ForgetPassword(db *gorm.DB, user *model.User) {

}

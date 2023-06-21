package controller

import (
	"gin/basic/dao"
	"gin/basic/db"
	"gin/basic/model"
	"gin/basic/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	//parse user data from body
	//call save user form dao
	user := model.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	err = dao.SaveUser(db.GetDBConn(), &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, user)
}
func LogIn(c *gin.Context) {
	/*
		1. parse user data from body
		2. match user details from db - dao.login
		3. generate token and send
	*/

	user := model.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	err = dao.LogIn(db.GetDBConn(), &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	//now you verified alreday that user is g[d i.e, exist

	token, err := utils.GenerateToekn(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, token)

}
func LogOut(c *gin.Context) {

}
func ForgetPassword(c *gin.Context) {

}

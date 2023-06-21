package router

import (
	"gin/basic/controller"

	"github.com/gin-gonic/gin"
)

func UserRouter(engine *gin.Engine) {
	engine.POST("/signup", controller.SignUp)
	engine.POST("/login", controller.LogIn)
	engine.POST("/logout", controller.LogOut)
	engine.PUT("/forgetpassword", controller.ForgetPassword)
}

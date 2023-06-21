package router

import (
	"gin/basic/controller"

	"github.com/gin-gonic/gin"
)

func PostRouter(engine *gin.Engine) {
	engine.GET("/", controller.HelloWorld)
	engine.GET("/post", controller.GetPost) //getPost this will come from controller
	engine.POST("/post", controller.CreatePost)
	engine.GET("/post/:id", controller.GetPostById)
	engine.PUT("/post", controller.UpdatePost)
	engine.DELETE("/post/:id", controller.DeletePostByID)
	//engine.PATCH("/post", controller.GetPost)
}

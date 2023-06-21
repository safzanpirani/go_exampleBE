package main

import (
	"fmt"
	"gin/basic/db"
	"gin/basic/router"
	"gin/basic/utils"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {

	checksum := func(c *gin.Context) {
		log.Println(c.Request.URL)
		if c.Request.URL.String() == "/login" || c.Request.URL.String() == "/signup" || c.Request.URL.String() == "/" {
			c.Next()
			return
		}
		token := c.GetHeader("token")
		isValid, err := utils.ValidateToken(token)
		if err != nil && !isValid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "Try again with new token")
		}
		c.Next()
	}
	return checksum
}

func main() {
	r := gin.Default()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	r.Use(AuthMiddleware())
	_, err := db.GetPostgresDBConnection()
	if err != nil {
		log.Fatal("Unable to establish DB connection", err)
	} else {
		log.Println("DB connection establsihed")
	}

	router.PostRouter(r)
	router.UserRouter(r)
	r.Run(fmt.Sprintf(":%s", port))

}

package controller

import (
	"fmt"
	"gin/basic/dao"
	"gin/basic/db"
	"gin/basic/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// acting as db for as
var post = map[string]string{
	"rahul": "Present",
}

func HelloWorld(c *gin.Context) {

	res := gin.H{
		"message": "Hello there!",
	}
	c.JSON(http.StatusOK, res)
}

func GetPost(c *gin.Context) {

	res := gin.H{
		"message": "",
	}
	listOfPosts, err := dao.GetPosts(db.GetDBConn())
	if err != nil {
		log.Print("got error", err)
		res = gin.H{
			"message": err.Error(),
		}
		c.JSON(http.StatusInternalServerError, res)
	} else {
		res = gin.H{
			"message": listOfPosts,
		}
		c.JSON(http.StatusOK, res)
	}
}
func DeletePostByID(c *gin.Context) {

	key := c.Params.ByName("id")
	res := gin.H{
		"message": "",
	}
	post, err := dao.GetPostByID(db.GetDBConn(), key)
	log.Println(fmt.Sprint(post.Id), "post.Id, key", len(fmt.Sprint(post.Id)), len(key), fmt.Sprint(post.Id) == key)
	if err != nil {
		log.Print("got error", err)
		res = gin.H{
			"message": err.Error(),
		}
		c.JSON(http.StatusInternalServerError, res)
	} else if fmt.Sprint(post.Id) == key { //means post exist
		post, err := dao.DeletePostById(db.GetDBConn(), key)
		if err != nil {

			log.Print("got error", err)
			res = gin.H{
				"message": err.Error(),
			}
			c.JSON(http.StatusInternalServerError, res)
		} else {
			res = gin.H{
				"post deleted": post,
			}
			c.JSON(http.StatusOK, res)
		}
	} else {
		res = gin.H{
			"post exists": false,
		}
		c.JSON(http.StatusNotFound, res)
	}
}

func CreatePost(c *gin.Context) {
	var p1 model.Post

	err := c.ShouldBindJSON(&p1)
	if err != nil {
		log.Print("got error", err)
	}
	res := gin.H{
		"message": "",
	}
	_, err = dao.CreatePost(db.GetDBConn(), &p1)
	if err != nil {
		log.Print("got error", err)
		res = gin.H{
			"message": err.Error(),
		}
		c.JSON(http.StatusInternalServerError, res)
	} else {
		res = gin.H{
			"message": "Done",
		}
		c.JSON(http.StatusOK, res)
	}

}

func GetPostById(c *gin.Context) {
	key := c.Params.ByName("id")
	res := gin.H{
		"message": "",
	}
	post, err := dao.GetPostByID(db.GetDBConn(), key)
	if err != nil {
		log.Print("got error", err)
		res = gin.H{
			"message": err.Error(),
		}
		c.JSON(http.StatusInternalServerError, res)
	} else {
		res = gin.H{
			"message": post,
		}
		c.JSON(http.StatusOK, res)
	}
}

func UpdatePost(c *gin.Context) {

	res := gin.H{
		"message": "",
	}
	var p1 model.Post

	err := c.ShouldBindJSON(&p1)
	if err != nil {
		log.Print("got error", err)
	}
	_, err = dao.UpdatePost(db.GetDBConn(), &p1)
	if err != nil {
		log.Print("got error", err)
		res = gin.H{
			"message": err.Error(),
		}
		c.JSON(http.StatusInternalServerError, res)
	} else {
		res = gin.H{
			"message": "Done",
		}
		c.JSON(http.StatusOK, res)
	}
}

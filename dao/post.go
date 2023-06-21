package dao

import (
	"gin/basic/model"
	"gorm.io/gorm"
)

func CreatePost(db *gorm.DB, post *model.Post) (*model.Post, error) {
	err := db.Save(post).Error
	return post, err
}

func GetPosts(db *gorm.DB) ([]model.Post, error) {
	listofPosts := []model.Post{}
	query := db.Select("posts.*")
	err := query.Find(&listofPosts).Error
	return listofPosts, err
}

func GetPostByID(db *gorm.DB, id string) (model.Post, error) {
	post := model.Post{}
	err := db.Select("posts.*").Where("posts.id", id).First(&post).Error

	return post, err
}
func DeletePosts(db *gorm.DB) ([]model.Post, error) {
	return nil, nil
}
func DeletePostById(db *gorm.DB, id string) (model.Post, error) {
	post := model.Post{}
	err := db.Where("id=?", id).Delete(&post).Error
	return post, err
}
func UpdatePost(db *gorm.DB, post *model.Post) (*model.Post, error) {
	err := db.Save(post).Error
	return post, err
}

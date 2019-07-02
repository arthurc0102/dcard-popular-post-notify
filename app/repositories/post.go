package repositories

import (
	"github.com/arthurc0102/dcard-popular-post-notify/app/db"
	"github.com/arthurc0102/dcard-popular-post-notify/app/models"
)

func GetPosts() (posts []models.Post) {
	db.Connection.Find(&posts)
	return
}

func CreatePost(post models.Post) error {
	return db.Connection.Create(&post).Error
}

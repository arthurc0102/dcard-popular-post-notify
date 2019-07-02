package models

// Post model
type Post struct {
	ID        int    `json:"id" gorm:"primary_key"`
	Title     string `json:"title" gorm:"-"`
	LikeCount int    `json:"likeCount" gorm:"-"`
}

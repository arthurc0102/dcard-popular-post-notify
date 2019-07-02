package services

import (
	"sort"

	"github.com/arthurc0102/dcard-popular-post-notify/app/repositories"

	"github.com/spf13/viper"

	"github.com/arthurc0102/dcard-popular-post-notify/app/models"
	"github.com/arthurc0102/dcard-popular-post-notify/helper"
	"github.com/parnurzeal/gorequest"
)

// GetPopularPosts return all popular posts
func GetPopularPosts(lessLikeCount int) ([]models.Post, []error) {
	var targetPosts []models.Post

	before := -1
	postsURL := viper.GetString("dcard.posts_url")

	for {
		var posts []models.Post

		query := map[string]interface{}{
			"popular": "true",
			"limit":   30,
		}

		if before != -1 {
			query["before"] = before
		}

		res, _, errs := gorequest.New().
			Get(postsURL).
			Query(query).
			EndStruct(&posts)

		if errs != nil {
			return nil, errs
		}

		helper.Log("(services/post) Send request to:", res.Request.URL)

		if AllPostsNotInTarget(posts, lessLikeCount) {
			break
		}

		posts = FilterTargetPosts(posts, lessLikeCount)
		targetPosts = append(targetPosts, posts...)
		before = posts[len(posts)-1].ID
	}

	sort.Slice(targetPosts, func(i int, j int) bool {
		return targetPosts[i].LikeCount > targetPosts[j].LikeCount
	})

	return targetPosts, nil
}

// AllPostsNotInTarget check if all post is not in target
func AllPostsNotInTarget(posts []models.Post, lessLikeCount int) bool {
	for _, post := range posts {
		if post.LikeCount >= lessLikeCount {
			return false
		}
	}

	return true
}

// FilterTargetPosts get posts in target
func FilterTargetPosts(posts []models.Post, lessLikeCount int) (targetPosts []models.Post) {
	for _, post := range posts {
		if post.LikeCount >= lessLikeCount {
			targetPosts = append(targetPosts, post)
		}
	}

	return
}

// FilterUnsentPosts return posts that haven't sent
func FilterUnsentPosts(posts []models.Post) (unsetPosts []models.Post) {
	sentPosts := make(map[int]string)

	for _, post := range repositories.GetPosts() {
		sentPosts[post.ID] = post.Title
	}

	for _, post := range posts {
		if _, ok := sentPosts[post.ID]; ok {
			continue
		}

		unsetPosts = append(unsetPosts, post)
	}

	return unsetPosts
}

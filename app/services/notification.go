package services

import (
	"fmt"

	"github.com/arthurc0102/dcard-popular-post-notify/helper"

	"github.com/arthurc0102/dcard-popular-post-notify/app/repositories"

	"github.com/parnurzeal/gorequest"

	"github.com/arthurc0102/dcard-popular-post-notify/app/models"
	"github.com/spf13/viper"
)

// SendToTGChannel send post to telegram channel
func SendToTGChannel(posts []models.Post) []error {
	postURL := viper.GetString("dcard.postURL")
	chatID := viper.GetString("telegram.chatID")
	disableWebPagePreview := viper.GetString("telegram.disableWebPagePreview")
	sendMessageURL := fmt.Sprintf(
		viper.GetString("telegram.sendMessageURL"),
		viper.GetString("telegram.token"),
	)

	posts = FilterUnsentPosts(posts)
	for _, post := range posts {
		link := fmt.Sprintf(postURL, post.ID)
		message := fmt.Sprintf("%s - %s", post.Title, link)
		query := map[string]interface{}{
			"chat_id":                  chatID,
			"disable_web_page_preview": disableWebPagePreview,
			"text":                     message,
		}

		_, _, errs := gorequest.New().
			Get(sendMessageURL).
			Query(query).
			End()

		if errs != nil {
			return errs
		}

		helper.Log("(services/notification) Notify:", message)

		err := repositories.CreatePost(post)
		if err != nil {
			return []error{err}
		}
	}

	return nil
}

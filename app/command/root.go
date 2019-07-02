package command

import (
	"os"

	"github.com/arthurc0102/dcard-popular-post-notify/app/db"

	"github.com/arthurc0102/dcard-popular-post-notify/app"
	"github.com/arthurc0102/dcard-popular-post-notify/app/services"
	"github.com/arthurc0102/dcard-popular-post-notify/helper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cmd = &cobra.Command{
	Use:               app.Name,
	PersistentPreRunE: helper.InitViper,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Load config
		lessLikeCount := viper.GetInt("lessLikeCount")
		dbConfig := viper.GetString("db.url")
		dbMigrate := viper.GetBool("db.migrate")

		// Database setup
		err := db.Setup(dbConfig)
		if err != nil {
			return err
		}

		defer func() {
			err := db.Connection.Close()
			if err != nil {
				panic(err)
			}
		}()

		if dbMigrate {
			db.Migrate()
		}

		// Do jobs
		posts, errs := services.GetPopularPosts(lessLikeCount)
		if errs != nil {
			return helper.MergeErrors(errs)
		}

		errs = services.SendToTGChannel(posts)
		if errs != nil {
			return helper.MergeErrors(errs)
		}

		return nil
	},
}

func init() {
	cmd.PersistentFlags().StringP("config", "c", "", "Config file.")
}

// Run command
func Run() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

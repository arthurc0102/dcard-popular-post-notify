package helper

import (
	"path/filepath"
	"strings"

	"github.com/arthurc0102/dcard-popular-post-notify/app"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// InitViper reads in config file and ENV variables if set.
func InitViper(cmd *cobra.Command, args []string) error {
	var configFile string

	if configFlag := cmd.Flags().Lookup("config"); configFlag != nil {
		configFile = configFlag.Value.String()
	}

	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		homeDirectory, err := HomeDirectory()
		if err != nil {
			return err
		}

		viper.SetConfigName("config")
		viper.SetConfigType("yaml")

		viper.AddConfigPath(".")
		viper.AddConfigPath(filepath.Join(homeDirectory, "."+app.Name))
		viper.AddConfigPath(filepath.Join("/", "etc", app.Name))
	}

	if err := bindFlags(cmd); err != nil {
		return err
	}

	viper.AutomaticEnv() // read in environment variables that match
	viper.SetEnvPrefix(app.ShortName)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	err := viper.ReadInConfig()

	switch err.(type) {
	case viper.ConfigFileNotFoundError:
		return nil // Ignore the error if config file is not found
	}

	Log("(helper/viper) Using config file:", viper.ConfigFileUsed())
	return err
}

func bindFlags(cmd *cobra.Command) error {
	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		return err
	}

	for _, subCmd := range cmd.Commands() {
		if err := bindFlags(subCmd); err != nil {
			return err
		}
	}

	return nil
}

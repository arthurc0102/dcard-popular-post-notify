package helper

import (
	"fmt"

	"github.com/spf13/viper"
)

// Log when is debug mode
func Log(message ...interface{}) {
	if !viper.GetBool("debug") {
		return
	}

	if len(message) == 0 {
		return
	}

	fmt.Println(message...)
}

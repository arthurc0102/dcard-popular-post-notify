package main

import (
	"fmt"
	"time"

	"github.com/arthurc0102/dcard-popular-post-notify/app/command"
)

func main() {
	command.Run()
	fmt.Println("Check finish at:", time.Now().Format("2006/01/02 15:04:05 -07:00"))
}

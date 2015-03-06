package main

import (
	"fmt"
	"github.com/patito/toby/conf"
)

func main() {

	fmt.Printf("Server = %s\n", conf.Config.Server)
	fmt.Printf("Nickname = %s\n", conf.Config.Nickname)
	fmt.Printf("Channel = %s\n", conf.Config.Channel)
	fmt.Printf("LogFile = %s\n", conf.LogFile)

	fmt.Printf("%p\n", conf.Config)

	return
}

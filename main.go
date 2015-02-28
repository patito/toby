package main

import (
	"fmt"
	conf "github.com/patito/toby/conf"
)

func main() {

	fmt.Printf("%s", conf.Get().Server)
	return
}

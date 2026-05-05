package main

import (
	"fmt"
	"goest/pkg/config"
)

func main() {

	config.Load()


	fmt.Println("Hello, World!")
	fmt.Println(config.Get())
}

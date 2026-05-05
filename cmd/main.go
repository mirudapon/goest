package main

import (
	"fmt"
	"goest/pkg/config"
)

type Config struct {
	A int32 `env:"A"`
}

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(config.Load[Config]())
}

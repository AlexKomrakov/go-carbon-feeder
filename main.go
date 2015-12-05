package main

import (
	"fmt"
)



func main() {


	config, err := Config{}.FromYml("config.yml")
	if err != nil {
		panic(err)
	}

	out, err := config.Command.Run()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Result: %s\n", out)
}
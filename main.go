package main

import (
	"./app"
	"./config"
	"fmt"
	"os"
)

func main() {
	configuration, err := config.Load("config.yml")
	if err != nil {
		fmt.Fprintf(os.Stdout, err.Error())
		return
	}

	app.Run(configuration)
}

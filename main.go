package main

import (
	"fmt"
	"github.com/albertogviana/port_scan/app"
	"github.com/albertogviana/port_scan/config"
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

package main

import (
	"github.com/malayanand/banking/app"
	"github.com/malayanand/banking/logger"
)

func main() {

	//fmt.Println("Starting the application...")
	logger.Info("Starting the application")
	app.Start()
}

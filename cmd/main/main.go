package main

import (
	"log"

	"github.com/IfuryI/WEB_BACK/internal/server"
	constants "github.com/IfuryI/WEB_BACK/pkg/const"
)

func main() {
	app := server.NewApp()

	if err := app.Run(constants.Port); err != nil {
		log.Fatal(err)
	}
}

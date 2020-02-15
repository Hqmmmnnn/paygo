package main

import (
	"log"

	"github.com/Hqqm/paygo/server"
	_ "github.com/lib/pq"
)

func main() {
	app := server.NewApp()

	if err := app.Run("8080"); err != nil {
		log.Fatal(err.Error())
	}
}

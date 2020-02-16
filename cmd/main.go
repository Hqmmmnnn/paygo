package main

import (
	"log"

	"github.com/Hqqm/paygo/config"
	"github.com/Hqqm/paygo/server"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.InitConfig("/paygo/config/config.yml")
	if err != nil {
		panic(err)
	}

	dsn := cfg.GetDsn()
	port := cfg.GetPort()

	app := server.NewApp(dsn)
	if err := app.Run(port); err != nil {
		log.Fatal(err.Error())
	}
}

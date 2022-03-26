package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"goo/db"
	"goo/routes"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")

	if read := viper.ReadInConfig(); read != nil {
		panic(read)
	}

	db.ConnectDb()

	app := fiber.New()

	port, host := viper.Get("server.port"), viper.Get("server.host")

	route.SetupRoute(app)
	app.Listen(fmt.Sprintf("%v:%v", host, port))
}

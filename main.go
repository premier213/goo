package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"goo/pkg/routes"
	"goo/platform/database"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./pkg/configs")

	if read := viper.ReadInConfig(); read != nil {
		panic(read)
	}

	database.ConnectDb()

	app := fiber.New()

	port, host := viper.Get("server.port"), viper.Get("server.host")

	route.SetupRoute(app)
	app.Listen(fmt.Sprintf("%v:%v", host, port))
}

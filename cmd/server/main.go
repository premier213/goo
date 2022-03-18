package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/premier213/goo/internal/routes"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./configs")

	if read := viper.ReadInConfig(); read != nil {
		panic(read)
	}

	app := fiber.New()

	port, host := viper.Get("server.port"), viper.Get("server.host")

	route.SetupRoute(app)
	app.Listen(fmt.Sprintf("%v:%v", host, port))
}

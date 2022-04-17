package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"goo/pkg/middleware"
	"goo/pkg/routes"
	"goo/platform/database"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./pkg/configs")
	// generate := uuid.New().String()

	// fmt.Println(string(generate[:5]))

	if read := viper.ReadInConfig(); read != nil {
		panic(read)
	}

	app := fiber.New()
	middleware.FiberMiddleware(app)

	route.SetupRoute(app)

	database.ConnectDb()

	port, host := viper.Get("server.port"), viper.Get("server.host")
	app.Listen(fmt.Sprintf("%v:%v", host, port))
}

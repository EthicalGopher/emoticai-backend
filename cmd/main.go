package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/EthicalGopher/emoticai/handlers"
	
)
func Startapp() {
	
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:	 "*",
		AllowMethods: 	"GET,POST,PUT,DELETE",
	}))
	app.Get("/",handlers.Homepage)
	defer app.Listen(":9045")

}

func main() {
	Startapp()
}
package main

import (
	"gopi/handlers"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	//Add config for connecting to CF IP
	app := fiber.New(fiber.Config{
		ProxyHeader: "CF-Connecting-IP",
		//EnableTrustedProxyCheck: true,
		//TrustedProxies: []string{"0.0.0.0"},
	})

	// Test route
	app.Get("/test", handlers.GetUserIPTest)

	// Main routes
	app.Get("/", handlers.GetCurrentIP)

	app.Get("/:ip/country", handlers.GetCountry)

	app.Get("/:ip", handlers.GetAnotherIP)

	log.Fatal(app.Listen(":3045"))
}

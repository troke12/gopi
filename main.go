package main

import (
	"github.com/troke12/gopi/handlers"
	"log"
	"github.com/joho/godotenv"
	"os"
	"github.com/gofiber/fiber/v2"
)

func main() {

	err := godotenv.Load()
  	if err != nil {
    	log.Fatal("Error loading .env file")
  	}
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

	log.Fatal(app.Listen(os.Getenv("PORT")))
}

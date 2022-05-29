package main

import (
	"github.com/troke12/gopi/handlers"
	"log"
	"github.com/joho/godotenv"
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/goccy/go-json"
	"github.com/rollbar/rollbar-go"
)

func main() {
	// Load environment
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Rollbar
	rollbar.SetToken(os.Getenv("ROLLBAR_TOKEN"))
	//Add config for connecting to CF IP
	app := fiber.New(fiber.Config{
		ProxyHeader: "CF-Connecting-IP",
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		//EnableTrustedProxyCheck: true,
		//TrustedProxies: []string{"0.0.0.0"},
	})

	// Test route
	//app.Get("/test", handlers.GetUserIPTest)

	// Main route
	app.Get("/", handlers.GetCurrentIP)
	rollbar.WrapAndWait(handlers.GetCurrentIP)

	// Getting country
	app.Get("/:ip/country", handlers.GetCountry)
	rollbar.WrapAndWait(handlers.GetCountry)

	// Request another IP
	app.Get("/:ip", handlers.GetAnotherIP)
	rollbar.WrapAndWait(handlers.GetAnotherIP)

	// Port
	log.Fatal(app.Listen(os.Getenv("PORT")))
}

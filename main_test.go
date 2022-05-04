// Testing units by https://dev.to/koddr/go-fiber-by-examples-testing-the-application-1ldf
package main

import (
	"net/http/httptest"
	"testing"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
		route string
		code int
	}{
		// First test
		{
			name: "get main route",
			route: "/",
			code: 200,
		},
		// Second test
		{
			name: "get http 404",
			route: "/not-found",
			code: 404,
		},
	}
	
	// Define fiber
	app := fiber.New()
	
	// Route for first test
	app.Get("/", func(c *fiber.Ctx) error {
		// return simple response
		return c.SendString("Hello!")
	})

	// main test
	for _, test := range tests {
		req := httptest.NewRequest("GET", test.route, nil)
		resp, _ := app.Test(req, 1)
		assert.Equalf(t, test.code, resp.StatusCode, test.name)
	}
}

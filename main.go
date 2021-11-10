package main

import (
	"fmt"
	"gopi/models"
	"gopi/handlers"
	"gopi/data"
	"log"
	"net"

	"github.com/gofiber/fiber/v2"
)

const DBIpGeo = "data/GeoIP2-City.mmdb"

func main() {
	//Add config for connecting to CF IP
	app := fiber.New(fiber.Config{
		ProxyHeader: "CF-Connecting-IP",
		//EnableTrustedProxyCheck: true,
		//TrustedProxies: []string{"0.0.0.0"},
	})

	//DBGet
	ConnectDB := dbmaxmind.GetDB(DBIpGeo)
	defer ConnectDB.Close()
	// Test route
	app.Get("/test", handlers.GetUserIP)

	// Main routes
	app.Get("/", func(c *fiber.Ctx) error {
		getclientIP := c.IP()
		ip := net.ParseIP(getclientIP)
		record, err := ConnectDB.City(ip)
		if err != nil {
			fmt.Printf("error: %v", err)
			return c.JSON(fiber.Map{
				"status": "err",
			})
		}
		//Logger
		fmt.Printf("%v\n", ip)
		dataIP := models.IpData{
			IP:            fmt.Sprintf("%v", ip),
			City:          record.City.Names["en"],
			Region:        record.City.Names["en"],
			Country:       record.Country.IsoCode,
			CountryFull:   record.Country.Names["en"],
			Continent:     record.Continent.Code,
			ContinentFull: record.Continent.Names["en"],
			Loc:           fmt.Sprintf("%v,%v", record.Location.Latitude, record.Location.Longitude),
			Postal:        record.Postal.Code,
		}
		return c.JSON(dataIP)
	})

	app.Get("/:ip/country", func(c *fiber.Ctx) error {
		ip := net.ParseIP(c.Params("ip"))
		record, err := ConnectDB.Country(ip)
		if err != nil {
			//log.Fatalf("error: %v", err)
			return c.JSON(fiber.Map{
				"status": "err",
			})
		}
		return c.SendString(record.Country.IsoCode)
	})

	app.Get("/:ip", func(c *fiber.Ctx) error {
		ip := net.ParseIP(c.Params("ip"))
		record, err := ConnectDB.City(ip)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			return c.JSON(fiber.Map{
				"status": "err",
			})
		}
		dataIP := models.IpData{
			IP:            fmt.Sprintf("%v", ip),
			City:          record.City.Names["en"],
			Region:        record.City.Names["en"],
			Country:       record.Country.IsoCode,
			CountryFull:   record.Country.Names["en"],
			Continent:     record.Continent.Code,
			ContinentFull: record.Continent.Names["en"],
			Loc:           fmt.Sprintf("%v,%v", record.Location.Latitude, record.Location.Longitude),
			Postal:        record.Postal.Code,
		}
		return c.JSON(dataIP)
	})

	log.Fatal(app.Listen(":3045"))
}

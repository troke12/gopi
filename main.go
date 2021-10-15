package main

import (
	"fmt"
	"log"
	"net"

	"github.com/gofiber/fiber/v2"
	"github.com/oschwald/geoip2-golang"
)

type IPstruct struct {
	IP            string
	City          string
	Region        string
	Country       string
	CountryFull   string
	Continent     string
	ContinentFull string
	Loc           string
	Postal        string
}

func main() {
	app := fiber.New()
	//Get the DB
	db, err := geoip2.Open("data/GeoIP2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		getClientIP := c.IP()
		ip := net.ParseIP(getClientIP)
		record, err := db.City(ip)
		if err != nil {
			log.Fatal(err)
		}
		//buat test di console
		//fmt.Printf("%v", record)

		dataIP := IPstruct{
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
		record, err := db.Country(ip)
		if err != nil {
			c.SendString("Error!")
		}
		return c.SendString(record.Country.IsoCode)
	})

	log.Fatal(app.Listen(":3000"))
}

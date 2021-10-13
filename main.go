package main

import (
	"fmt"
	"log"
	"net"

	"github.com/gofiber/fiber/v2"
	"github.com/oschwald/geoip2-golang"
)

type IPstruct struct {
	IP            string `json:"ip"`
	City          string `json:"city"`
	Region        string `json:"region"`
	Country       string `json:"country"`
	CountryFull   string `json:"country_full"`
	Continent     string `json:"continent"`
	ContinentFull string `json:"continent_full"`
	Loc           string `json:"loc"`
	Postal        string `json:"postal"`
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		db, err := geoip2.Open("data/GeoIP2-City.mmdb")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		getClientIP := c.IP()
		ipnya := net.ParseIP(getClientIP)
		record, err := db.City(ipnya)
		if err != nil {
			log.Fatal(err)
		}
		//buat test di console
		//fmt.Printf("%v", record)

		locationfull := fmt.Sprintf("%v,%v", record.Location.Latitude, record.Location.Longitude)
		ambilIP := fmt.Sprintf("%v", ipnya)
		dataIP := IPstruct{
			IP:            ambilIP,
			City:          record.City.Names["en"],
			Region:        record.City.Names["en"],
			Country:       record.Country.IsoCode,
			CountryFull:   record.Country.Names["en"],
			Continent:     record.Continent.Code,
			ContinentFull: record.Continent.Names["en"],
			Loc:           locationfull,
			Postal:        record.Postal.Code,
		}
		return c.JSON(dataIP)
	})

	log.Fatal(app.Listen(":3000"))
}

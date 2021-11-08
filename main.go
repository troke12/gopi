package main

import (
	"fmt"
	"log"
	"net"
	"strings"

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

//Add Test

func getUserIP(c *fiber.Ctx) error {
   	var userIP string
   	if len(c.GetRespHeader("CF-Connecting-IP")) > 1 {
        	userIP = c.GetRespHeader("CF-Connecting-IP")
        	fmt.Println(net.ParseIP(userIP))
    	} else if len(c.GetRespHeader("X-Forwarded-For")) > 1 {
        	userIP = c.GetRespHeader("X-Forwarded-For")
        	fmt.Println(net.ParseIP(userIP))
    	} else if len(c.GetRespHeader("X-Real-IP")) > 1 {
        	userIP = c.GetRespHeader("X-Real-IP")
        	fmt.Println(net.ParseIP(userIP))
    	} else {
        	userIP = c.IP()
        	if strings.Contains(userIP, ":") {
            		fmt.Println(net.ParseIP(strings.Split(userIP, ":")[0]))
        	} else {
            		fmt.Println(net.ParseIP(userIP))
        	}
 	}
	return c.JSON(fiber.Map{
		"ip": net.ParseIP(userIP),
	})
}

func main() {
	//Add config for connecting to CF IP
	app := fiber.New(fiber.Config{
		ProxyHeader: "CF-Connecting-IP",
	})
	//Get the DB
	db, err := geoip2.Open("data/GeoIP2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// Routes
	app.Get("/jancok", getUserIP)
	app.Get("/test", func(c *fiber.Ctx) error {
		h := c.IP()
		fmt.Println(h)
		return c.JSON(fiber.Map{
			"ip": h,
		})
	})
	app.Get("/", func(c *fiber.Ctx) error {
		getclientIP := c.IP()
		ip := net.ParseIP(getclientIP)
		record, err := db.City(ip)
		if err != nil {
			fmt.Printf("error: %v", err)
			return c.JSON(fiber.Map{
				"status": "err",
			})
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
			//log.Fatalf("error: %v", err)
			return c.JSON(fiber.Map{
				"status": "err",
			})
		}
		return c.SendString(record.Country.IsoCode)
	})

	app.Get("/:ip", func(c *fiber.Ctx) error {
		ip := net.ParseIP(c.Params("ip"))
		record, err := db.City(ip)
		if err != nil {
			//log.Fatalf("error: %v", err)
			return c.JSON(fiber.Map{
				"status": "err",
			})
		}
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

	log.Fatal(app.Listen(":3045"))
}

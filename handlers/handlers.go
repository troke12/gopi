package handlers

import (
	"fmt"
	"net"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"github.com/troke12/gopi/data"
	"github.com/troke12/gopi/models/local"
	"github.com/troke12/gopi/models/web"
	"strings"
	"os"

	"github.com/gofiber/fiber/v2"
)

// DB maxmind location
var DBIpGeo = "data/GeoIP2-City.mmdb"

// This test function for checking Cloudflare or another Header
func GetUserIPTest(c *fiber.Ctx) error {
	var userIP string
	if len(c.GetRespHeader("CF-Connecting-IP")) > 1 {
		userIP = c.GetRespHeader("CF-Connecting-IP")
		fmt.Println("Cloudflare check", net.ParseIP(userIP))
	} else if len(c.GetRespHeader("X-Forwarded-For")) > 1 {
		userIP = c.GetRespHeader("X-Forwarded-For")
		fmt.Println("Forwarded for", net.ParseIP(userIP))
	} else if len(c.GetRespHeader("X-Real-IP")) > 1 {
		userIP = c.GetRespHeader("X-Real-IP")
		fmt.Println("Real IP", net.ParseIP(userIP))
	} else {
		userIP = c.IP()
		if strings.Contains(userIP, ":") {
			fmt.Println("Natural IP6", net.ParseIP(strings.Split(userIP, ":")[0]))
		} else {
			fmt.Println("Natural IP", net.ParseIP(userIP))
		}
	}
	return c.JSON(fiber.Map{
		"ip": net.ParseIP(userIP),
	})
}

func GetCurrentIP(c *fiber.Ctx) error {
	getclientIP := c.IP()
	ip := net.ParseIP(getclientIP)
	record, err := dbmaxmind.GetDB(DBIpGeo).City(ip)
	if err != nil {
		fmt.Printf("error: %v", err)
		return c.JSON(fiber.Map{
			"status": "err",
		})
	}
	//Logger
	fmt.Printf("%v\n", ip)
	dataIP := local.IpData{
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
}

func GetCountry(c *fiber.Ctx) error {
	if strings.Contains(c.Params("ip"), ":") {
		var ipAddress string = c.Params("ip")
		v6 := net.ParseIP(ipAddress)
		baseUrl := fmt.Sprintf("https://api.freegeoip.app/json/%v?apikey=%v", v6, os.Getenv("API_KEY"))

		res, err := http.Get(baseUrl)
		if err != nil {
			fmt.Println("error cuy")
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)

		if err != nil {
			fmt.Println("No response from request")
		}

		var dataFree web.IpFG
		if err := json.Unmarshal(body, &dataFree); err != nil {
			fmt.Println("Error")
		}
		webIP := local.IpData{
			IP:				fmt.Sprintf("%v", ipAddress),
			City:          	dataFree.City,
			Region:        	dataFree.RegionName,
			Country:       	dataFree.CountryCode,
			CountryFull:   	dataFree.CountryName,
			Continent:     	dataFree.RegionCode,
			ContinentFull: 	dataFree.RegionName,
			Loc:           	fmt.Sprintf("%v,%v", dataFree.Latitude, dataFree.Longitude),
			Postal:        	dataFree.ZipCode,
		}
		return c.SendString(webIP.Country)
	} else {
		ip := net.ParseIP(c.Params("ip"))
		record, err := dbmaxmind.GetDB(DBIpGeo).Country(ip)
		if err != nil {
			//log.Fatalf("error: %v", err)
			return c.JSON(fiber.Map{
				"status": "err",
			})
		}
		return c.SendString(record.Country.IsoCode)
	}
}

func GetAnotherIP(c *fiber.Ctx) error {
	if strings.Contains(c.Params("ip"), ":") {
		var ipAddress string = c.Params("ip")
		v6 := net.ParseIP(ipAddress)
		baseUrl := fmt.Sprintf("https://api.freegeoip.app/json/%v?apikey=%v", v6, os.Getenv("API_KEY"))

		res, err := http.Get(baseUrl)
		if err != nil {
			fmt.Println("error cuy")
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)

		if err != nil {
			fmt.Println("No response from request")
		}

		var dataFree web.IpFG
		if err := json.Unmarshal(body, &dataFree); err != nil {
			fmt.Println("Error")
		}
		webIP := local.IpData{
			IP:				fmt.Sprintf("%v", ipAddress),
			City:          	dataFree.City,
			Region:        	dataFree.RegionName,
			Country:       	dataFree.CountryCode,
			CountryFull:   	dataFree.CountryName,
			Continent:     	dataFree.RegionCode,
			ContinentFull: 	dataFree.RegionName,
			Loc:           	fmt.Sprintf("%v,%v", dataFree.Latitude, dataFree.Longitude),
			Postal:        	dataFree.ZipCode,
		}
		return c.JSON(webIP)
	} else {
		ipv4 := net.ParseIP(c.Params("ip"))
		record, err := dbmaxmind.GetDB(DBIpGeo).City(ipv4)
		if err != nil {
			//fmt.Printf("error: %v\n", err)
			return c.JSON(fiber.Map{
				"status": "err",
			})
		}
		dataIP := local.IpData{
			IP:            fmt.Sprintf("%v", ipv4),
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
	}
}
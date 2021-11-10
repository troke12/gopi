package handlers

import (
	"fmt"
	"net"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GetUserIP(c *fiber.Ctx) error {
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

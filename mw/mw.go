package mw

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func Auth(c *fiber.Ctx) error {
	log.Println("mw log")
	return c.Next()
}

package router

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

var Server *fiber.App

func Initialize() {
	Server = fiber.New()
}

func Start() {
	log.Fatal(Server.Listen(":3000"))
}

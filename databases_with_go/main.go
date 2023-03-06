package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", indexHandler)     // Index handler
	app.Post("/", postHandler)     // Post Handler
	app.Put("/update", putHandler) // Put Handlerr
	app.Delete("/delete", deleteHandler)

	port := os.Getenv("PORT")
	fmt.Println(port)
	if port == "" {
		port = "3000"
	}
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}

func indexHandler(c *fiber.Ctx) error {
	return c.SendString("Hello")
}

func postHandler(c *fiber.Ctx) error {
	return c.SendString("Hello")
}

func putHandler(c *fiber.Ctx) error {
	return c.SendString("Hello")
}

func deleteHandler(c *fiber.Ctx) error {
	return c.SendString("Hello")
}

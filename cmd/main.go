package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kwesikwaa/toyshop-backend/internal/handlers"
)

func firsfunc(c *fiber.Ctx) error {
	return c.SendString("Yh firs one of the shop api")
}

func main() {
	app := fiber.New()

	app.Use("/api", func(c *fiber.Ctx) error {
		fmt.Println("middleware middle ground")
		return c.Next()
	})

	app.Get("/", firsfunc)

	app.Get("api/v1/products", handlers.GetAllProducts)
	// app.Get("api/v1/products/:id",handlers.GetSingleProduct)
	app.Post("api/v1/products", handlers.CreateProduct)

	fmt.Println("something")

	log.Fatal(app.Listen(":3000"))
}

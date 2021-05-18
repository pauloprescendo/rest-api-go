package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pauloprescendo/go-rest-api/pauloprescendo/rest-api-go/database"
	"github.com/pauloprescendo/go-rest-api/pauloprescendo/rest-api-go/router"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	database.ConnectDB()

	router.SetupRoutes(app)

	app.Listen(":3000")
}

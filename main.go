package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/qinains/fastergoding"

	"github.com/guhkun13/tutorial/fiber-api/database"
	"github.com/guhkun13/tutorial/fiber-api/routes"
)

func main() {
	app := fiber.New()
	fastergoding.Run()

	// set middleware
	app.Use(logger.New())

	// set database
	database.ConnectDb()
	log.Println("Database migrated!")

	// set routes
	routes.SetupRoutes(app)

	// start server
	port := ":8000"
	log.Printf("Start server at port %s \n", port)
	log.Fatal(app.Listen(port))
}
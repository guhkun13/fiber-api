package routes

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App){
	// welcome endpoint 
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Halo dunia")
	})
	
	userApp := app.Group("users")
	userApp.Get("/", GetUsers)
	userApp.Post("/", CreateUser)
	userApp.Get("/:id", GetUser)
	userApp.Put("/:id", UpdateUser)
	userApp.Delete("/:id", DeleteUser)
}
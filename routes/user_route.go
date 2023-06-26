package routes

import (
	"github.com/Alex980102/go-mongo-fiber/controllers"
	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	// Create a user
	app.Post("/user", controllers.CreateUser)

	// Get user by ID
	app.Get("/user/:userId", controllers.GetAUser)

	// Edit a user
	app.Put("/user/:userId", controllers.EditAUser)

	// Delete a user
	app.Delete("/user/:userId", controllers.DeleteAUser)

	// Get all users
	app.Get("/users", controllers.GetAllUsers) //add this
}

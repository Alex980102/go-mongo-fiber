package main

import (
	"crypto/sha256"
	"crypto/subtle"

	"github.com/Alex980102/go-mongo-fiber/configs"
	"github.com/Alex980102/go-mongo-fiber/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var (
	apiKey = "correct horse battery staple"
)

func validateAPIKey(c *fiber.Ctx, key string) (bool, error) {
	hashedAPIKey := sha256.Sum256([]byte(apiKey))
	hashedKey := sha256.Sum256([]byte(key))

	if subtle.ConstantTimeCompare(hashedAPIKey[:], hashedKey[:]) == 1 {
		return true, nil
	}
	return false, keyauth.ErrMissingOrMalformedAPIKey
}

func main() {
	app := fiber.New()

	// note that the keyauth middleware needs to be defined before the routes are defined!
	app.Use(keyauth.New(keyauth.Config{
		KeyLookup: "cookie:access_token",
		Validator: validateAPIKey,
	}))

	// Add Cors
	app.Use(cors.New())

	app.Use(logger.New())

	// Connect to MongoDB
	configs.ConnectDB()

	// Routes
	routes.UserRoute(app)

	app.Listen(":6000")
}

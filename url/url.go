package url

import (
	"github.com/nekowawolf/crypto-community-api/controller"

	"github.com/gofiber/fiber/v2"
)

func Web(page *fiber.App) {

	page.Get("/", controller.Sink)
	page.Post("/", controller.Sink)
	page.Put("/", controller.Sink)
	page.Patch("/", controller.Sink)
	page.Delete("/", controller.Sink)
	page.Options("/", controller.Sink)

	cryptoGroup := page.Group("/cryptocommunity")
	cryptoGroup.Get("/", controller.GetAllCryptoCommunity)
	cryptoGroup.Get("/:id", controller.GetCryptoCommunityByID)
	cryptoGroup.Get("/search/:name", controller.GetCryptoCommunityByName)
	cryptoGroup.Post("/", controller.InsertCryptoCommunity)
	cryptoGroup.Put("/:id", controller.UpdateCryptoCommunityByID)
	cryptoGroup.Delete("/:id", controller.DeleteCryptoCommunityByID)
	
}
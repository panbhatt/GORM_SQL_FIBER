package router

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/panbhatt/gorm_fiber_sqlserver/docs"
	noteRoutes "github.com/panbhatt/gorm_fiber_sqlserver/internals/routes/note"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/swagger/*", swagger.Handler)

	// app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
	// 	URL:         "http://example.com/doc.json",
	// 	DeepLinking: false,
	// 	// Expand ("list") or Collapse ("none") tag groups by default
	// 	DocExpansion: "none",
	// 	// Prefill OAuth ClientId on Authorize popup
	// 	OAuth: &swagger.OAuthConfig{
	// 		AppName:  "OAuth Provider",
	// 		ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
	// 	},
	// 	// Ability to change OAuth2 redirect uri location
	// 	OAuth2RedirectUrl: "http://localhost:8080/swagger/oauth2-redirect.html",
	// }))
	// Group api calls with param '/api'
	api := app.Group("/api", logger.New())

	// Setup note routes, can use same syntax to add routes for more models
	noteRoutes.SetupNoteRoutes(api)
}

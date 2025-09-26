
package api_routes

import (
	"github.com/MarcelArt/learn-react-go/database"
	api_handlers "github.com/MarcelArt/learn-react-go/handlers/api"
	"github.com/MarcelArt/learn-react-go/middlewares"
	"github.com/MarcelArt/learn-react-go/repositories"
	"github.com/gofiber/fiber/v2"
)

func SetupRoleRoutes(api fiber.Router, auth *middlewares.AuthMiddleware) {
	h := api_handlers.NewRoleHandler(repositories.NewRoleRepo(database.GetDB()))

	g := api.Group("/role")
	g.Get("/", auth.ProtectedAPI, h.Read)
	g.Get("/:id", auth.ProtectedAPI, h.GetByID)
	g.Post("/", auth.ProtectedAPI, h.Create)
	g.Put("/:id", auth.ProtectedAPI, h.Update)
	g.Delete("/:id", auth.ProtectedAPI, h.Delete)
}

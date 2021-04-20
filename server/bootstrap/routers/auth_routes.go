package routers

import (
	"dew-backend/server/handlers"

	"github.com/gofiber/fiber/v2"
)

// AuthRoutes ...
type AuthRoutes struct {
	RouterGroup fiber.Router
	Handler     handlers.Handler
}

// RegisterRoute register auth routes
func (route AuthRoutes) RegisterRoute() {
	handler := handlers.AuthHandler{Handler: route.Handler}

	r := route.RouterGroup.Group("/api/auth")
	r.Post("/login", handler.Login)
}

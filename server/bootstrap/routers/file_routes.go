package routers

import (
	"dew-backend/server/handlers"
	"dew-backend/server/middlewares"

	"github.com/gofiber/fiber/v2"
)

// FileRoutes ...
type FileRoutes struct {
	RouterGroup fiber.Router
	Handler     handlers.Handler
}

// RegisterRoute register user routes
func (route FileRoutes) RegisterRoute() {
	handler := handlers.FileHandler{Handler: route.Handler}
	jwtMiddleware := middlewares.JwtMiddleware{ContractUC: handler.ContractUC}

	user := route.RouterGroup.Group("/api/file")
	user.Use(jwtMiddleware.VerifyUser)
	user.Post("", handler.Upload)
}

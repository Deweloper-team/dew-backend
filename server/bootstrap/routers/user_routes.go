package routers

import (
	"tradesignal-backend/server/handlers"
	"tradesignal-backend/server/middlewares"

	"github.com/gofiber/fiber/v2"
)

// UserRoutes ...
type UserRoutes struct {
	RouterGroup fiber.Router
	Handler     handlers.Handler
}

// RegisterRoute register user routes
func (route UserRoutes) RegisterRoute() {
	handler := handlers.UserHandler{Handler: route.Handler}
	jwtMiddleware := middlewares.JwtMiddleware{ContractUC: handler.ContractUC}

	user := route.RouterGroup.Group("/api/user")
	user.Use(jwtMiddleware.VerifyUser)
	user.Get("", handler.Get)

	test := route.RouterGroup.Group("/api/test")
	test.Use(jwtMiddleware.VerifyBasic)
	test.Get("/mail", handler.TestMail)
}

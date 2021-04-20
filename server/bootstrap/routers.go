package bootstrap

import (
	"net/http"
	"tradesignal-backend/server/bootstrap/routers"
	"tradesignal-backend/server/handlers"

	"github.com/gofiber/fiber/v2"
)

// RegisterRouters ...
func (boot Bootstrap) RegisterRouters() {
	handler := handlers.Handler{
		FiberApp:   boot.App,
		ContractUC: &boot.ContractUC,
		Validator:  boot.Validator,
		Translator: boot.Translator,
	}

	// Testing
	boot.App.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON("work")
	})

	apiV1 := boot.App.Group("/v1")

	// auth routes
	authRoutes := routers.AuthRoutes{RouterGroup: apiV1, Handler: handler}
	authRoutes.RegisterRoute()

	// user routes
	userRoutes := routers.UserRoutes{RouterGroup: apiV1, Handler: handler}
	userRoutes.RegisterRoute()

	// file routes
	fileRoutes := routers.FileRoutes{RouterGroup: apiV1, Handler: handler}
	fileRoutes.RegisterRoute()
}

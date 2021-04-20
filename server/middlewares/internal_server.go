package middlewares

import (
	"tradesignal-backend/pkg/functioncaller"
	"tradesignal-backend/pkg/logruslogger"
	"tradesignal-backend/usecase/viewmodel"

	"github.com/gofiber/fiber/v2"
)

func InternalServer(ctx *fiber.Ctx, err error) error {
	// Statuscode defaults to 500
	code := fiber.StatusInternalServerError

	// Retreive the custom statuscode if it's an fiber.*Error
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	logruslogger.Log(logruslogger.ErrorLevel, err.Error(), functioncaller.PrintFuncName(), "internal_server")
	return ctx.Status(code).JSON([]interface{}{viewmodel.ResponseErrorVM{Messages: err.Error()}})
}

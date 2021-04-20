package handlers

import (
	"context"
	"dew-backend/pkg/str"
	"dew-backend/server/requests"
	"dew-backend/usecase"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// AuthHandler ...
type AuthHandler struct {
	Handler
}

// Login ...
func (h *AuthHandler) Login(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(context.Background(), time.Duration(str.StringToInt(h.ContractUC.EnvConfig["APP_TIMEOUT"]))*time.Second)
	c = context.WithValue(c, "requestid", ctx.Locals("requestid"))
	defer cancel()

	input := new(requests.LoginRequest)
	if err := ctx.BodyParser(input); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}
	input.RemoteIP = ctx.IP()
	if err := h.Validator.Struct(input); err != nil {
		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
	}

	authUc := usecase.AuthUC{ContractUC: h.ContractUC}
	res, err := authUc.Login(c, input)

	return h.SendResponse(ctx, res, nil, err, 0)
}

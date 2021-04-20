package handlers

import (
	"context"
	"dew-backend/helper"
	"dew-backend/pkg/str"
	"dew-backend/usecase"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

// UserHandler ...
type UserHandler struct {
	Handler
}

// Get ...
func (h *UserHandler) Get(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(context.Background(), time.Duration(str.StringToInt(h.ContractUC.EnvConfig["APP_TIMEOUT"]))*time.Second)
	c = context.WithValue(c, "requestid", ctx.Locals("requestid"))
	c = context.WithValue(c, "user_id", ctx.Locals("user_id"))
	defer cancel()

	user := h.GetUser(ctx)
	clientID := ctx.Query("client_id")
	if clientID == "" {
		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
	}

	uc := usecase.UserUC{ContractUC: h.ContractUC}
	res, err := uc.Get(c, clientID, &user)

	return h.SendResponse(ctx, res, nil, err, 0)
}

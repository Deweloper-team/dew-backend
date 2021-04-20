package handlers

import (
	"context"
	"net/http"
	"time"
	"tradesignal-backend/helper"
	"tradesignal-backend/pkg/str"
	"tradesignal-backend/usecase"

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

// TestMail ...
func (h *UserHandler) TestMail(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(context.Background(), time.Duration(str.StringToInt(h.ContractUC.EnvConfig["APP_TIMEOUT"]))*time.Second)
	c = context.WithValue(c, "requestid", ctx.Locals("requestid"))
	c = context.WithValue(c, "user_id", ctx.Locals("user_id"))
	defer cancel()

	email := ctx.Query("email")
	subject := ctx.Query("subject")
	body := ctx.Query("body")

	uc := usecase.UserUC{ContractUC: h.ContractUC}
	err := uc.TestMail(c, email, subject, body)

	return h.SendResponse(ctx, nil, nil, err, 0)
}

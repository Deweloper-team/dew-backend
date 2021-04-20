package handlers

import (
	"context"
	"errors"
	"time"
	"tradesignal-backend/db/repository/models"
	"tradesignal-backend/helper"
	"tradesignal-backend/pkg/str"
	"tradesignal-backend/usecase"

	"github.com/gofiber/fiber/v2"
)

// FileHandler ...
type FileHandler struct {
	Handler
}

// Upload ...
func (h *FileHandler) Upload(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(context.Background(), time.Duration(str.StringToInt(h.ContractUC.EnvConfig["APP_TIMEOUT"]))*time.Second)
	c = context.WithValue(c, "requestid", ctx.Locals("requestid"))
	c = context.WithValue(c, "user_id", ctx.Locals("user_id"))
	defer cancel()

	// Read file type
	fileType := ctx.FormValue("type")
	if !str.Contains(models.FileWhitelist, fileType) {
		return h.SendResponse(ctx, nil, nil, errors.New(helper.InvalidFileType), 0)
	}

	// Upload file to local temporary
	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		return h.SendResponse(ctx, nil, nil, err, 0)
	}
	fileUc := usecase.FileUC{ContractUC: h.ContractUC}
	res, err := fileUc.Upload(c, fileType, fileHeader)

	return h.SendResponse(ctx, res, nil, err, 0)
}

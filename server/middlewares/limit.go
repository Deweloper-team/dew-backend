package middlewares

import (
	"dew-backend/usecase"
)

// LimitInit ...
type LimitInit struct {
	*usecase.ContractUC
	MaxLimit float64
	Duration string
}

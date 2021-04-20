package middlewares

import (
	"tradesignal-backend/usecase"
)

// LimitInit ...
type LimitInit struct {
	*usecase.ContractUC
	MaxLimit float64
	Duration string
}

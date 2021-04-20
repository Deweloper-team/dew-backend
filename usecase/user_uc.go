package usecase

import (
	"context"
	"database/sql"
	"dew-backend/usecase/viewmodel"
)

// UserUC ...
type UserUC struct {
	ContractUC *ContractUC
	Tx         *sql.Tx
}

// Get ...
func (uc UserUC) Get(c context.Context, clientID string, data *viewmodel.UserVM) (res viewmodel.UserVM, err error) {
	return res, err
}

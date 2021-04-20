package usecase

import (
	"context"
	"database/sql"
	"dew-backend/pkg/functioncaller"
	"dew-backend/pkg/logruslogger"
	"dew-backend/server/requests"
	"dew-backend/usecase/viewmodel"
)

// AuthUC ...
type AuthUC struct {
	ContractUC *ContractUC
	Tx         *sql.Tx
}

// Login ...
func (uc AuthUC) Login(c context.Context, data *requests.LoginRequest) (res viewmodel.JwtVM, err error) {
	// Decrypt password input
	data.Password, err = uc.ContractUC.AesFront.Decrypt(data.Password)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "decrypt", c.Value("requestid"))
		return res, err
	}

	// Jwe the payload & Generate jwt token
	payload := map[string]interface{}{
		"user_id": data.User,
	}

	jwtUc := JwtUC{ContractUC: uc.ContractUC}
	err = jwtUc.GenerateToken(c, payload, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "generate_token", c.Value("requestid"))
		return res, err
	}
	res.UserID = data.User

	return res, err
}

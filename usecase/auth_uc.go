package usecase

import (
	"context"
	"database/sql"
	"tradesignal-backend/pkg/functioncaller"
	"tradesignal-backend/pkg/logruslogger"
	"tradesignal-backend/pkg/str"
	"tradesignal-backend/server/requests"
	"tradesignal-backend/usecase/viewmodel"
)

// AuthUC ...
type AuthUC struct {
	ContractUC *ContractUC
	Tx         *sql.Tx
}

// Login ...
func (uc AuthUC) Login(c context.Context, data *requests.LoginRequest) (res viewmodel.JwtVM, err error) {
	// Run recaptcha validation
	if str.StringToBool(uc.ContractUC.EnvConfig["APP_USE_RECAPTCHA"]) {
		recaptchaUc := RecaptchaUC{ContractUC: uc.ContractUC}
		_, err = recaptchaUc.Verify(c, data.Recaptcha, data.RemoteIP)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "recaptcha", c.Value("requestid"))
			return res, err
		}
	}

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

package usecase

import (
	"context"
	"database/sql"
	"tradesignal-backend/pkg/functioncaller"
	"tradesignal-backend/pkg/logruslogger"
	"tradesignal-backend/usecase/viewmodel"
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

// TestMail ...
func (uc UserUC) TestMail(c context.Context, email, subject, body string) (err error) {
	mailUc := MailUC{ContractUC: uc.ContractUC}
	err = mailUc.Send(c, email, subject, body)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "send_mail", c.Value("requestid"))
		return err
	}

	return err
}

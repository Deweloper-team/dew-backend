package usecase

import (
	"context"

	"tradesignal-backend/pkg/logruslogger"
)

// MailUC ...
type MailUC struct {
	*ContractUC
}

// Send ...
func (uc MailUC) Send(c context.Context, to, subject, body string) (err error) {
	ctx := "MailUC.Send"

	if uc.EnvConfig["SMTP_PROVIDER"] == "mandrill" {
		err = uc.ContractUC.Mailing.Send(to, subject, body)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "mandrill", c.Value("requestid"))
			return err
		}
	} else {
		err = uc.ContractUC.Mail.Send(uc.EnvConfig["SMTP_FROM"], []string{to}, subject, body)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "smtp", c.Value("requestid"))
			return err
		}
	}

	return err
}

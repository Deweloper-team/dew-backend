package google

import (
	"github.com/appleboy/go-fcm"
)

type Credential struct {
	Key string
}

func (cred Credential) SendFcm(payload map[string]interface{}, fcmToken []string, title, body string) error {
	msg := &fcm.Message{
		RegistrationIDs: fcmToken,
		Notification: &fcm.Notification{
			Title: title,
			Body:  body,
		},
		Data: payload,
	}

	client, err := fcm.NewClient(cred.Key)
	if err != nil {
		return err
	}

	_, err = client.Send(msg)
	if err != nil {
		return err
	}

	return nil
}

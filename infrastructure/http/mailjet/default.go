package mailjet

import (
	"context"
	"fmt"
	"os"

	"github.com/fadilahonespot/cakrawala/utils/logger"
	"github.com/mailjet/mailjet-apiv3-go"
)

type wrapper struct {
	client *mailjet.Client
}

func NewWrapper() MailjetWrapper {
	mailjetClient := mailjet.NewMailjetClient(os.Getenv("MAILJET_PUBLIC_KEY"), os.Getenv("MAILJET_PRIVATE_KEY"))

	return &wrapper{client: mailjetClient}
}

func (w *wrapper) SendEmail(ctx context.Context, email, title, body string) (err error) {
	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: os.Getenv("MAILJET_FROM_EMAIL"),
				Name:  os.Getenv("MAILJET_FROM_NAME"),
			},
			To: &mailjet.RecipientsV31{
				{
					Email: email,
				},
			},
			Subject:  title,
			HTMLPart: body,
		},
	}

	messages := mailjet.MessagesV31{Info: messagesInfo}
	logger.Info(ctx, "[SendEmail Request]", messages)

	resp, err := w.client.SendMailV31(&messages)
	if err != nil {
		return fmt.Errorf("failed to send email %v", err.Error())
	}

	logger.Info(ctx, "[Send Email Response]", resp)

	if len(resp.ResultsV31) != 0 {
		if resp.ResultsV31[0].Status != StatusSuccess {
			return fmt.Errorf("status %v is not success", resp.ResultsV31[0].Status)
		}
	}
	return
}

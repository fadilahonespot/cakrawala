package mailjet

import (
	"context"
)

type MailjetWrapper interface {
	SendEmail(ctx context.Context, email, title, body string) (err error)
}
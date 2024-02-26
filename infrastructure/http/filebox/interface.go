package filebox

import (
	"context"
	"mime/multipart"
)

type DropboxWrapper interface {
	Uplaod(ctx context.Context, file *multipart.FileHeader, pathFolder string) (link string, err error)
}
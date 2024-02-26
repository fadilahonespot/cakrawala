package filebox

import (
	"context"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/fadilahonespot/cakrawala/utils"
	"github.com/tj/go-dropbox"
	"github.com/tj/go-dropy"
)

type wrapper struct {
	client *dropy.Client
}

func NewWrapper() DropboxWrapper {
	token := os.Getenv("DROPBOX_ACCESS_TOKEN")
	client := dropy.New(dropbox.New(dropbox.NewConfig(token)))
	return &wrapper{client: client}
}

func (s *wrapper) Uplaod(ctx context.Context, file *multipart.FileHeader, pathFolder string) (link string, err error) {
	pathFile := pathFolder + fmt.Sprintf("file_%v", utils.GenerateRandomString(20)) + filepath.Ext(file.Filename)
	src, err := file.Open()
	if err != nil {
		err = fmt.Errorf("error opening file %v: %v", pathFolder, err)
		return
	}
	defer src.Close()
	
	err = s.client.Upload(pathFile, src)
	if err != nil {
		err = fmt.Errorf("error uploading file %v", err.Error())
		return
	}
	share := new(dropbox.CreateSharedLinkInput)
	share.Path = pathFile
	sharedLink, err := s.client.Sharing.CreateSharedLink(share)
	if err != nil {
		err = fmt.Errorf("error create share link: %v", err.Error())
		return
	}
	sharedLink.Path = pathFile

	link = cutStringFile(file.Filename, sharedLink.URL)
	return 
}

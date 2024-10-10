package supabase

import (
	"mime/multipart"
	"net/url"
	"time"

	supabaseService "github.com/adityarizkyramadhan/supabase-storage-uploader"
)

type Supabase interface {
	Upload(fileHeader *multipart.FileHeader) (string, error)
	Delete(link string) error
}

type supabase struct {
	Client *supabaseService.Client
}

func New(projectURL string, token string, bucketName string) Supabase {
	client := supabaseService.New(projectURL, token, bucketName)
	return &supabase{client}
}

func (s *supabase) Upload(fileHeader *multipart.FileHeader) (string, error) {
	fileHeader.Filename = url.QueryEscape(time.Now().String() + fileHeader.Filename)
	return s.Client.Upload(fileHeader)
}

func (s *supabase) Delete(link string) error {
	return s.Client.Delete(link)
}
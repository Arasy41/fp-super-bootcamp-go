package utils

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"path"
	"time"

	"api-culinary-review/config"

	supabasestorageuploader "github.com/adityarizkyramadhan/supabase-storage-uploader"
	storage_go "github.com/supabase-community/storage-go"
)

func UploadFileToSupabase(file *multipart.FileHeader) (string, error) {
	conf := config.LoadConfig()

	client := storage_go.NewClient(conf.SupabaseURL, conf.SupabaseKey, nil)
	bucketName := conf.SupabaseBucket

	fileName := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
	filePath := path.Join("public", fileName)

	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer src.Close()

	buffer := make([]byte, file.Size)
	if _, err = io.ReadFull(src, buffer); err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	reader := bytes.NewReader(buffer)

	res, err := client.UploadFile(bucketName, filePath, reader)
	if err != nil {
		return "", fmt.Errorf("failed to upload file: %w", err)
	}

	if res.Error != "" {
		return "", fmt.Errorf("upload error: %v", res.Error)
	}

	url := fmt.Sprintf("https://%s.storage.supabase.co/storage/v1/object/%s/%s", conf.SupabaseURL, bucketName, filePath)
	return url, nil
}

func DeleteImageFromSupabase(imagePath string) error {
	conf := config.LoadConfig()

	client := supabasestorageuploader.New(
		conf.SupabaseURL,
		conf.SupabaseKey,
		conf.SupabaseBucket,
	)

	err := client.Delete(imagePath)
	if err != nil {
		return fmt.Errorf("failed to delete image: %v", err)
	}

	return nil
}

package utils

import (
	"api-culinary-review/config"
	"context"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadToCloudinary(file *multipart.FileHeader) (string, error) {
	// Buat direktori assets/uploads/ jika belum ada
	if _, err := os.Stat("internal/assets/uploads/"); os.IsNotExist(err) {
		err = os.MkdirAll("internal/assets/uploads/", os.ModePerm)
		if err != nil {
			return "", err
		}
	}

	// Simpan file sementara di assets/uploads/
	tempFilePath := filepath.Join("internal/assets/uploads/", file.Filename)
	out, err := os.Create(tempFilePath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	in, err := file.Open()
	if err != nil {
		return "", err
	}
	defer in.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return "", err
	}

	// Remove file setelah upload ke Cloudinary
	defer func() {
		os.Remove(tempFilePath)
	}()

	cloudinaryURL := config.LoadConfig().CloudinaryURL
	cld, err := cloudinary.NewFromURL(cloudinaryURL)
	if err != nil {
		return "", err
	}

	var ctx = context.Background()
	uid, err := GenerateUid()
	if err != nil {
		return "", err
	}
	resp, err := cld.Upload.Upload(ctx, tempFilePath, uploader.UploadParams{PublicID: "image-" + file.Filename + "-" + uid})
	if err != nil {
		return "", err
	}

	return resp.SecureURL, nil
}

func DeleteImageFromCloudinary(imageURL string) error {
	cloudinaryURL := config.LoadConfig().CloudinaryURL
	cld, err := cloudinary.NewFromURL(cloudinaryURL)
	if err != nil {
		return err
	}

	var ctx = context.Background()
	_, err = cld.Upload.Destroy(ctx, uploader.DestroyParams{PublicID: imageURL})
	if err != nil {
		return err
	}

	return nil
}

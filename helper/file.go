package helper

import (
	"io"
	"mime/multipart"
	"os"
)

const (
	FilePath = "/files"
)

func SaveFile(filename string, file *multipart.FileHeader) (string, error) {
	src, err := file.Open()

	if err != nil {
		return "", err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(FilePath + filename)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return "", err
	}

	return FilePath + filename, nil
}

func RemoveFile(filepath string) {
	os.Remove(filepath)
}

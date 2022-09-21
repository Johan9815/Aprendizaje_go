package utils

import (
	"io"
	"mime/multipart"
	"os"
)

func UploadImage(file *multipart.FileHeader) string {

	src, err := file.Open()
	if err != nil {
		panic(err)
	}
	defer src.Close()

	// Destination
	dst, err := os.Create("src/public/images/" + file.Filename)
	if err != nil {
		panic(err)
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		panic(err)
	}

	return file.Filename
}

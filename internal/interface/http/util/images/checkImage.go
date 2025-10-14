package images

import (
	"errors"
	"mime/multipart"
	"net/http"
)

func CheckImage(file *multipart.FileHeader, allowTypes map[string]bool, maxSize int) error {
	// چک کردن سایز
	if err := checkSize(file, maxSize); err != nil {
		return err
	}

	// باز کردن فایل
	src, err := file.Open()
	defer func(src multipart.File) {
		err := src.Close()
		if err != nil {
			return
		}
	}(src)
	if err != nil {
		return err
	}

	// چک کردن فرمت
	if err := checkFormat(src, allowTypes); err != nil {
		return err
	}

	return nil
}

func checkSize(file *multipart.FileHeader, maxSize int) error {
	if file.Size > int64(maxSize) {
		err := errors.New("file size too big")
		return err
	}
	return nil
}
func checkFormat(src multipart.File, allowTypes map[string]bool) error {
	buffer := make([]byte, 512)
	_, err := src.Read(buffer)
	if err != nil {
		return err
	}
	mimeType := http.DetectContentType(buffer)
	if _, ok := allowTypes[mimeType]; !ok {
		return errors.New("File type not allowed: " + mimeType)
	}
	return nil
}

package images

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"mime/multipart"
	"net/http"
	"path/filepath"
)

func CreateImage(c *fiber.Ctx, allowTypes map[string]bool, maxSize int) (newPhotoDbPath string, err error) {
	// گرفتن فایل
	teammatePhoto, err := c.FormFile("image")
	newPhotoDbPath = ""
	if err == nil {
		if err := CheckImage(teammatePhoto, allowTypes, maxSize); err != nil {
			return newPhotoDbPath, errors.New("image type is not allow or image size is too big") // ارور نا مناسب
		}

		// ساخت اسم جدید
		fileName := uuid.New().String() + filepath.Ext(teammatePhoto.Filename)
		savePath := filepath.Join("./images", fileName)
		newPhotoDbPath = "/images/" + fileName

		if err := CrateFile(c, teammatePhoto, savePath); err != nil {
			return newPhotoDbPath, errors.New("can't save image " + fileName + ": " + err.Error())
		}
	} else if !errors.Is(err, http.ErrMissingFile) {
		return newPhotoDbPath, errors.New("you have to enter a file")
	}

	return newPhotoDbPath, nil
}

func CrateFile(c *fiber.Ctx, teammatePhoto *multipart.FileHeader, savePath string) error {
	err := c.SaveFile(teammatePhoto, savePath)
	if err != nil {
		return err
	}
	return nil
}

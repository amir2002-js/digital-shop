package images

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
)

func ActionImg(newArr, oldArr *pq.StringArray, c *fiber.Ctx) error {
	action := c.Query("img_action", "replace")

	if *newArr != nil {
		switch action {
		case "append":
			exists := make(map[string]bool)
			for _, img := range *oldArr {
				exists[img] = true
			}
			for _, img := range *newArr {
				if !exists[img] {
					*oldArr = append(*oldArr, img)
				}
			}

		case "remove":
			removeSet := make(map[string]bool)
			for _, img := range *newArr {
				removeSet[img] = true
			}
			var filtered pq.StringArray
			for _, img := range *oldArr {
				if !removeSet[img] {
					filtered = append(filtered, img)
				}
			}
			*oldArr = filtered
		case "replace":
			*oldArr = *newArr
		default:
			return errors.New("invalid img_action: must be one of [append, remove, replace]")
		}
	}

	return nil
}

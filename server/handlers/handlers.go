package handlers

import (
	"go-gorm-mk1-showcase/gorm"

	"github.com/gofiber/fiber/v2"
)

// HANDLERS

func Error(c *fiber.Ctx, err error) error {
	switch err {
	case fiber.ErrNotFound:
		return c.SendString("Page doesn't exists: 404. " + err.Error())

	default:
		return c.SendString("Something bad happened. " + err.Error())
	}
}

func Root(c *fiber.Ctx) error {
	err := c.Render("index", fiber.Map{})
	if err != nil {
		return err
	}
	return nil
}

func Query(c *fiber.Ctx) error {
	query := c.Params("query")

	switch query {
	case "select-all":
		all := gorm.SelectAll()

		err := c.JSON(all)
		if err != nil {
			return err
		}

	case "select-where":
		where := gorm.SelectWhere()

		err := c.JSON(where)
		if err != nil {
			return err
		}

	case "select-specific":
		specific := gorm.SelectSpecific()

		err := c.JSON(specific)
		if err != nil {
			return err
		}

	case "update-all":
		updated := gorm.UpdateAll()
		err := c.JSON(updated)
		if err != nil {
			return err
		}

	case "update-name":
		updated := gorm.UpdateName()

		err := c.JSON(updated)
		if err != nil {
			return err
		}

	case "delete-row":
		affected := gorm.DeleteRow()

		err := c.JSON(affected)
		if err != nil {
			return err
		}

	case "insert-row":
		inserted := gorm.InsertRow()

		err := c.JSON(inserted)
		if err != nil {
			return err
		}
	}

	return nil
}

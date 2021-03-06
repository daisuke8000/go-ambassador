package controllers

import (
	"ambassador-backend/src/database"
	"ambassador-backend/src/models"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func Link(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var links []models.Link

	database.DB.Where("user_id=?", id).Find(&links)

	return c.JSON(links)
}
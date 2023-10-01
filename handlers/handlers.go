package handlers

import (
	"eventsWebFiber/db"
	"eventsWebFiber/models"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func Health(c *fiber.Ctx) error {
	return c.SendString("Hello")
}

func CreateEvent(c *fiber.Ctx) error {
	var webEvent models.WebEvent

	dbConnection, err := db.GetConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError)
		return err
	}

	if dbConnection == nil {
		c.JSON(http.StatusInternalServerError)
		return err
	}

	if conversionError := c.BodyParser(&webEvent); conversionError != nil {
		c.JSON(http.StatusInternalServerError)
		return conversionError
	}
	dbResult := dbConnection.Connection.Create(&webEvent)
	if dbResult.Error != nil {
		c.JSON(http.StatusInternalServerError)
		return dbResult.Error
	}
	c.JSON(webEvent)
	return nil
}

func ListEvents(c *fiber.Ctx) error {
	dbConnection, err := db.GetConnection()
	if err != nil || dbConnection == nil {
		c.JSON(http.StatusInternalServerError)
		return err
	}

	var webEvents []models.WebEvent
	dbResult := dbConnection.Connection.Find(&webEvents)

	if dbResult.Error != nil {
		c.JSON(http.StatusInternalServerError)
		return dbResult.Error
	}

	c.JSON(webEvents)
	return nil
}

func FindEvent(c *fiber.Ctx) error {
	dbConnection, err := db.GetConnection()
	if err != nil || dbConnection == nil {
		c.JSON(http.StatusInternalServerError)
		return err
	}
	var event models.WebEvent
	id := c.Params("id")
	dbResult := dbConnection.Connection.First(&event, id)
	if dbResult.Error != nil {
		c.JSON(http.StatusNotFound)
		return dbResult.Error
	}
	c.JSON(event)
	return nil
}

func DeleteEvent(c *fiber.Ctx) error {
	dbConnection, err := db.GetConnection()
	if err != nil || dbConnection == nil {
		c.JSON(http.StatusInternalServerError)
		return err
	}
	id := c.Params("id")
	dbResult := dbConnection.Connection.Delete(&models.WebEvent{}, id)
	if dbResult.Error != nil {
		c.JSON(http.StatusInternalServerError)
		return dbResult.Error
	}
	c.JSON(http.StatusOK)
	return nil
}

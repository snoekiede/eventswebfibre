package main

import (
	"eventsWebFiber/handlers"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	app.Get("/", handlers.Health)
	app.Get("/events", handlers.ListEvents)
	app.Get("/event/:id", handlers.FindEvent)
	app.Post("/event/create", handlers.CreateEvent)
	app.Delete("/event/delete/:id", handlers.DeleteEvent)

	log.Fatal(app.Listen(":8080"))
}

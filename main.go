package main

import (
	"go-gorm-mk1-showcase/gorm"
	"go-gorm-mk1-showcase/server/handlers"

	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

// ENTRY POINT

func main() {
	err := gorm.ConfigGorm()
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New(fiber.Config{
		Views:        html.New("./server/templates", ".html"),
		ReadTimeout:  time.Second * 20,
		WriteTimeout: time.Second * 20,
		ErrorHandler: handlers.Error,
		AppName:      "go-gorm-mk1-showcase",
	})

	app.Get("/", handlers.Root)
	app.Get("/:query", handlers.Query)

	err = app.Listen("0.0.0.0:3240")
	if err != nil {
		log.Fatal(err)
	}
}

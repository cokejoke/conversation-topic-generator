package main

import (
	"conversation-topic/config"
	"conversation-topic/database"
	"conversation-topic/resource"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	config.InitConfiguration()

	database.Connect()
	defer database.DBConn.Close()

	app := fiber.New()
	app.Use(cors.New())

	resources := loadResources()
	for _, element := range resources {
		element.Init(app)
	}

	app.Listen(":8991")
}

func loadResources() []resource.Resource {
	var resources []resource.Resource
	resources = append(resources, resource.TopicResource{}, resource.PackResource{}, resource.AuthResource{})
	return resources
}

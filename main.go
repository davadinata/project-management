package main

import (
	"log"

	"github.com/davadinata/project-management/config"
	"github.com/davadinata/project-management/controller"
	"github.com/davadinata/project-management/database/seed"
	"github.com/davadinata/project-management/repositories"
	"github.com/davadinata/project-management/routes"
	"github.com/davadinata/project-management/services"
	"github.com/gofiber/fiber/v3"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	seed.SeedAdmin()

	app := fiber.New()

	userRepo := repositories.NewUserRepository()

	userService := services.NewUserService(userRepo)

	userController := controller.NewUserController(userService)

	routes.Setup(app, userController)

	port := config.AppConfig.APPPort

	log.Println("Server is running on port", port)
	log.Fatal(app.Listen(":" + port))
}

package seed

import (
	"log"
	"os"

	"github.com/davadinata/project-management/config"
	"github.com/davadinata/project-management/models"
	"github.com/davadinata/project-management/utils"
)

func SeedAdmin() {
	password, _ := utils.HashPassword("admin123")

	admin := models.User{
		Name:     "Super Admin",
		Email:    os.Getenv("ADMIN_EMAIL"),
		Password: password,
		Role:     os.Getenv("ADMIN_ROLE"),
	}

	if err := config.DB.FirstOrCreate(&admin, models.User{Email: "admin.Email"}).Error; err != nil {
		log.Println("failed to seed Admin.", err)
	} else {
		log.Println("Admin seeded successfully.")
	}
}

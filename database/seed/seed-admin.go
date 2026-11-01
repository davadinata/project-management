package seed

import (
	"log"
	"os"

	"github.com/davadinata/project-management/config"
	"github.com/davadinata/project-management/models"
	"github.com/davadinata/project-management/utils"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func SeedAdmin() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Gagal load .env:", err)
	}
	password, err := utils.HashPassword(os.Getenv("ADMIN_PASSWORD"))
	if err != nil {
		log.Fatal("Gagal hash password:", err)
	}
	admin := models.User{
		Name:     "Super Admin",
		Email:    os.Getenv("ADMIN_EMAIL"),
		Password: password,
		Role:     os.Getenv("ADMIN_ROLE"),
		PublicID: uuid.New(),
	}

	if err := config.DB.FirstOrCreate(&admin, models.User{Email: admin.Email}).Error; err != nil {
		log.Println("failed to seed Admin.", err)
	} else {
		log.Println("Admin seeded successfully.")
	}
}

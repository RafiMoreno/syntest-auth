package initializers

import (
	"auth-service/models"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {
	DB.AutoMigrate(&models.User{})
}
package migrations

import (
	"github.com/brunohs007/CRUD/tree/develop/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Usuario{})
}

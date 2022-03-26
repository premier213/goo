package migrations

import (
	"fmt"
	"goo/pkg/models"
	"log"

	"gorm.io/gorm"
)

var model = []interface{}{
	&models.User{},
}

func AutoMigrateDB(DB *gorm.DB) error {
	log.Println("Running migrations")

	if err := DB.AutoMigrate(model...); err != nil {
		fmt.Println(err)
	}
	return nil
}

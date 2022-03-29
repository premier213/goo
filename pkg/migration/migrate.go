package migration

import (
	"fmt"
	"goo/pkg/model"
	"log"

	"gorm.io/gorm"
)

var models = []interface{}{
	&model.User{},
}

func AutoMigrateDB(DB *gorm.DB) error {
	log.Println("Running migrations")

	if err := DB.AutoMigrate(models...); err != nil {
		fmt.Println(err)
	}
	return nil
}

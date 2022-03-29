package db

import (
	"fmt"
	"github.com/spf13/viper"
	"goo/pkg/migration"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

type DbInstance struct {
	Db *gorm.DB
}

var DB DbInstance

func ConnectDb() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("../../config/config.yml")

	host, port, user, password, dbname, sslmode, timezone := viper.Get("database.host"), viper.Get("database.port"), viper.Get("database.user"), viper.Get("database.password"), viper.Get("database.dbName"), viper.Get("database.sslMode"), viper.Get("database.timezone")

	if read := viper.ReadInConfig(); read != nil {
		panic(read)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s", host, user, password, dbname, port, sslmode, timezone)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	log.Println("connected to database 🔥")
	db.Logger = logger.Default.LogMode(logger.Info)

	migration.AutoMigrateDB(db)

	DB = DbInstance{
		Db: db,
	}

}

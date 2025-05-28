package database

import (
	"fmt"
	"log"

	"github.com/fajardwntara/vow-connect/api/config"
	"github.com/fajardwntara/vow-connect/api/domain/user"
	"github.com/fajardwntara/vow-connect/api/domain/wedding"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(cfg *config.Config) {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	// DB.Migrator().DropTable(&user.User{})

	if err := Migrate(DB); err != nil {
		log.Fatal("failed to migrate the models of database:", err)
	}

	log.Println("*** Database connected and migrated succesfully ***")
}

func GetDB() *gorm.DB {
	return DB
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&user.User{},
		&user.Organizer{},
		&wedding.Wedding{},
		&user.Guest{},
	)
}

package db

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/Danil-114195722/TemplateFilesManager/db/models"
	"github.com/Danil-114195722/TemplateFilesManager/settings"
)


// создание таблиц в БД по структурам в Go
func Migrate() {
	fmt.Println("Start migration...")
	
	fmt.Printf("Connectiong to %q database...\n", settings.PathDB)
	dbConnect := GetConnection()

	err := dbConnect.AutoMigrate(&models.File{})
	settings.DieIf(err)

	fmt.Println("DB -- Migrated successfully!")
}

// получение соединения с БД
func GetConnection() *gorm.DB {
	connection, err := gorm.Open(sqlite.Open(settings.PathDB), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	settings.DieIf(err)

	return connection
}

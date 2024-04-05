package initializers

import (
	"fmt"
	"gorm.io/driver/mysql"
	"os"

	"github.com/robbyklein/pages/models"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SyncDB() {
	DB.AutoMigrate(&models.Person{})
}

func ConnectToDB() {
	var err error
	DB, err = gorm.Open(mysql.Open(os.Getenv("DB")), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed to connect to db")
	}
}

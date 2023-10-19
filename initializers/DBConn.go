package initializers

import (

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	dsn := "host=floppy.db.elephantsql.com user=sfnnxfvv password=2qIhhGm00kWQjsp3TKXwJ8wn_Xysjgze dbname=sfnnxfvv port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to DB")
	}
}

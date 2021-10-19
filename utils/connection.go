package utils

import (
	"log"

	"github.com/jinzhu/gorm"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

// GetConnection obtiene una conexión a la base de datos
func GetConnection() *gorm.DB {

	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

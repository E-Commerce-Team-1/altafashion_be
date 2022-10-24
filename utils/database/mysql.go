package database

import (
	"altafashion_be/config"
	rUser "altafashion_be/feature/users/repository"

	"fmt"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(c *config.AppConfig) *gorm.DB {
	str := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.DBUser,
		c.DBPwd,
		c.DBHost,
		c.DBPort,
		c.DBName,
	)

	db, err := gorm.Open(mysql.Open(str), &gorm.Config{})
	if err != nil {
		log.Error("db config error :", err.Error())
		return nil
	}
	migrateDB(db)
	return db
}

// func InitDB(cfg *config.AppConfig) *gorm.DB {
// 	dsn := "root:@tcp(localhost:3307)/medsos?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		log.Fatal("Cannot connect to DB")
// 	}

// 	migrateDB(db)
// 	return db
// }

func migrateDB(db *gorm.DB) {
	db.AutoMigrate(&rUser.User{})
}

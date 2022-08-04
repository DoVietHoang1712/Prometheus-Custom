package database

import (
	"PrometheusCustom/model"
	"PrometheusCustom/util"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDb() *gorm.DB {
	Db = connectDB()
	return Db
}

func connectDB() *gorm.DB {
	var err error
	config, _ := util.LoadConfig(".")
	dsn := config.DBUsername + ":" + config.DBPassword + "@tcp" + "(" + config.DBHost + ":" + config.DBPort + ")/" + config.DBName + "?" + "parseTime=true&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to database : error=%v", err)
		return nil
	}

	db.AutoMigrate(&model.CpuOversaturion{})
	db.AutoMigrate(&model.PodRestarted{})
	return db
}

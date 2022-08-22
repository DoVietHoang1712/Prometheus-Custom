package database

import (
	"PrometheusCustom/model"
	"PrometheusCustom/util"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDb() *gorm.DB {
	Db = connectDB()
	return Db
}

func connectDB() *gorm.DB {
	var err error
	config, _ := util.LoadConfig()
	dsn := "postgres://" + config.DBUsername + ":" + config.DBPassword + "@" + config.DBHost + ":" + config.DBPort + "/" + config.DBName
	fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to database : error=%v", err)
		return nil
	}

	db.AutoMigrate(&model.CpuOversaturion{})
	db.AutoMigrate(&model.PodRestarted{})
	db.AutoMigrate(&model.MemoryUnderUtilize{})
	return db
}

package main

import (
	"PrometheusCustom/database"
	"PrometheusCustom/model"
)

func main() {
	//now := time.Now()
	//config, _ := util.LoadConfig()
	//hour, minute := util.GetTimeStart(config.TimeStart)
	//loc, _ := time.LoadLocation(config.TimeZone)
	//t := time.Date(now.Year(), now.Month(), now.Day(), int(hour), int(minute), 0, 0, loc)
	//gocron.Every(1).Day().From(&t).Do(func() {
	db := database.InitDb()
	db.AutoMigrate(&model.CpuOversaturion{})
	db.AutoMigrate(&model.PodRestarted{})
	cpuOversaturationResponse := GetCpuOversaturation()
	model.CreateCpuOversaturion(db, cpuOversaturationResponse)
	podRestartedResponse := GetPodRestarted()
	model.CreatePodRestarted(db, podRestartedResponse)
	//})
	//<-gocron.Start()
}

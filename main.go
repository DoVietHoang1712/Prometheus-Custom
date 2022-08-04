package main

import (
	"PrometheusCustom/database"
	"PrometheusCustom/model"
	"time"
)

func main() {
	db := database.InitDb()
	db.AutoMigrate(&model.CpuOversaturion{})
	db.AutoMigrate(&model.PodRestarted{})
	for {
		cpuOversaturationResponse := GetCpuOversaturation()
		model.CreateCpuOversaturion(db, cpuOversaturationResponse)
		podRestartedResponse := GetPodRestarted()
		model.CreatePodRestarted(db, podRestartedResponse)
		time.Sleep(15 * time.Minute)
	}

}

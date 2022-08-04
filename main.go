package main

import (
	"PrometheusCustom/database"
	"PrometheusCustom/model"
)

func main() {
	db := database.InitDb()
	db.AutoMigrate(&model.CpuOversaturion{})
	db.AutoMigrate(&model.PodRestarted{})
	cpuOversaturationResponse := GetCpuOversaturation()
	model.CreateCpuOversaturion(db, cpuOversaturationResponse)
	podRestartedResponse := GetPodRestarted()
	model.CreatePodRestarted(db, podRestartedResponse)
}

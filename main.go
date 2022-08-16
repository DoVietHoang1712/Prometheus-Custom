package main

import (
	"PrometheusCustom/database"
	"PrometheusCustom/model"
	"github.com/jasonlvhit/gocron"
)

func main() {

	gocron.Every(1).Day().At("17:13").Do(func() {
		db := database.InitDb()
		db.AutoMigrate(&model.CpuOversaturion{})
		db.AutoMigrate(&model.PodRestarted{})
		cpuOversaturationResponse := GetCpuOversaturation()
		model.CreateCpuOversaturion(db, cpuOversaturationResponse)
		podRestartedResponse := GetPodRestarted()
		model.CreatePodRestarted(db, podRestartedResponse)
	})
	<-gocron.Start()
}

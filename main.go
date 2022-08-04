package main

import (
	"PrometheusCustom/database"
	"PrometheusCustom/model"
)

type Metric struct {
	Cluster   string `json:"cluster"`
	Instance  string `json:"instance"`
	Job       string `json:"job"`
	Namespace string `json:"namespace"`
	Pod       string `json:"pod"`
}

type Result struct {
	Metric Metric        `json:"metric"`
	Value  []interface{} `json:"value"`
}

type Data struct {
	ResultType string   `json:"resultType"`
	Result     []Result `json:"result"`
}

type Response struct {
	Status string `json:"status"`
	Data   Data   `json:"data"`
}

func main() {
	db := database.InitDb()
	db.AutoMigrate(&model.CpuOversaturion{})
	db.AutoMigrate(&model.PodRestarted{})
	cpuOversaturationResponse := GetCpuOversaturation()
	model.CreateCpuOversaturion(db, cpuOversaturationResponse)
	podRestartedResponse := GetPodRestarted()
	model.CreatePodRestarted(db, podRestartedResponse)
}

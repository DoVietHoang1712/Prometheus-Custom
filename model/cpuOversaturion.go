package model

import (
	"PrometheusCustom/util"
	"gorm.io/gorm"
)

type CpuOversaturion struct {
	ID                uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Workload          string `gorm:"size:255;not null" json:"workload"`
	Cluster           string `gorm:"size:255;not null;" json:"cluster"`
	Namespace         string `gorm:"size:255;null;" json:"namespace"`
	Container         string `gorm:"size:255;null" json:"container"`
	CpuSaturation     float64
	SuggestCpuRequest float64
	Time              int64
	WorkloadInfo      util.MetricWorkload `gorm:"-"`
}

type CpuOversaturionResponse struct {
	ID                uint32 `json:"id"`
	Workload          string `json:"workload"`
	Cluster           string `json:"cluster"`
	Namespace         string `json:"namespace"`
	Container         string `json:"container"`
	CpuSaturation     float64
	SuggestCpuRequest float64
	Value             float64
	Time              int64
	WorkloadInfo      util.MetricWorkload
}

func CreateCpuOversaturion(db *gorm.DB, CpuOversaturion []CpuOversaturion) (err error) {
	err = db.Create(&CpuOversaturion).Error
	if err != nil {
		return err
	}
	return nil
}

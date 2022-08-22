package model

import (
	"PrometheusCustom/util"
	"gorm.io/gorm"
)

type MemoryUnderUtilize struct {
	ID                     uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Workload               string `gorm:"size:255;not null" json:"workload"`
	Cluster                string `gorm:"size:255;not null;" json:"cluster"`
	Namespace              string `gorm:"size:255;null;" json:"namespace"`
	Container              string `gorm:"size:255;null" json:"container"`
	MemoryUnderUtilize1Day float64
	MemoryUnderUtilize3Day float64
	MemoryUnderUtilize7Day float64
	SuggestMemoryLimit     int64
	SuggestMemoryRequest   int64
	Time                   int64
	WorkloadInfo           util.MetricWorkload `gorm:"-"`
}

type MemoryUnderUtilizeResponse struct {
	ID                     uint32 `json:"id"`
	Workload               string `json:"workload"`
	Cluster                string `json:"cluster"`
	Namespace              string `json:"namespace"`
	Container              string `json:"container"`
	MemoryUnderUtilize1Day float64
	MemoryUnderUtilize3Day float64
	MemoryUnderUtilize7Day float64
	SuggestMemoryLimit     float64
	SuggestMemoryRequest   int64
	Value                  float64
	Time                   int64
	WorkloadInfo           util.MetricWorkload
}

func CreateMemoryUnderUtilize(db *gorm.DB, MemoryUnderUtilize []MemoryUnderUtilize) (err error) {
	err = db.Create(&MemoryUnderUtilize).Error
	if err != nil {
		return err
	}
	return nil
}

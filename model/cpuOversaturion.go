package model

import "gorm.io/gorm"

type CpuOversaturion struct {
	ID                uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Workload          string `gorm:"size:255;not null" json:"workload"`
	Cluster           string `gorm:"size:255;not null;" json:"cluster"`
	SuggestCpuRequest float64
	Time              int64
	WorkloadInfo      util.MetricWorkload `gorm:"-"`
}

func CreateCpuOversaturion(db *gorm.DB, CpuOversaturion []CpuOversaturion) (err error) {
	err = db.Create(&CpuOversaturion).Error
	if err != nil {
		return err
	}
	return nil
}

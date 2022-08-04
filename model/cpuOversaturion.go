package model

import "gorm.io/gorm"

type CpuOversaturion struct {
	ID      uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Pod     string `gorm:"size:255;not null" json:"pod"`
	Cluster string `gorm:"size:255;not null;" json:"cluster"`
	Time    int64
}

func CreateCpuOversaturion(db *gorm.DB, CpuOversaturion []CpuOversaturion) (err error) {
	err = db.Create(&CpuOversaturion).Error
	if err != nil {
		return err
	}
	return nil
}

package model

import "gorm.io/gorm"

type PodRestarted struct {
	ID      uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Pod     string `gorm:"size:255;not null" json:"pod"`
	Cluster string `gorm:"size:255;not null;" json:"cluster"`
	Time    int64
}

func CreatePodRestarted(db *gorm.DB, PodRestarted []PodRestarted) (err error) {
	err = db.Create(&PodRestarted).Error
	if err != nil {
		return err
	}
	return nil
}

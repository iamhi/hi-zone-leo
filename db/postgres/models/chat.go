package models

import "gorm.io/gorm"

type Chat struct {
	gorm.Model
	Uuid      string `gorm:"unique" json:"uuid"`
	OwnerUuid string `json:"owner_uuid"`
}

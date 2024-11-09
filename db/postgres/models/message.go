package models

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	Uuid     string `gorm:"unique" json:"uuid"`
	ChatUuid string `json:"chat_uuid"`
	Role     string `json:"role"`
	Content  string `json:"content"`
	Visible  bool   `json:"visible"`
}

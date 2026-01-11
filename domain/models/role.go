package models

import "time"

type Role struct {
	ID   uint    `gorm:"primaryKey;autoIncrement"`
	Code string  `gorm:"varchar(100);not null"`
	Name string  `gorm:"varchar(100);not null"`
	CreatedAt *time.Time  `gorm:"autoCreateTime"`
	UpdatedAt *time.Time  `gorm:"autoUpdateTime"`
}

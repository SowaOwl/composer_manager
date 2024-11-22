package model

import "time"

type PublicType struct {
	ID         uint   `gorm:"primary_key"`
	Name       string `gorm:"unique;not null"`
	ManyName   string
	IsTabulate bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

package model

import "time"

type Network struct {
	ID         uint   `gorm:"primary_key"`
	Name       string `gorm:"unique"`
	Body       string `gorm:"type:text"`
	ComposerID uint   `gorm:"index"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

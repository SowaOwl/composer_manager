package model

import "time"

type Composer struct {
	ID        uint   `gorm:"primary_key"`
	Path      string `gorm:"unique"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Containers []Container `gorm:"foreignKey:ComposerID"`
	Networks   []Network   `gorm:"foreignKey:ComposerID"`
}

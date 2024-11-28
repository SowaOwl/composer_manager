package model

import "time"

type PublicData struct {
	ID           uint `gorm:"primary_key"`
	PublicTypeID uint `gorm:"index;not null"`
	ContainerID  uint `gorm:"index;not null"`
	Data         string
	CreatedAt    time.Time
	UpdatedAt    time.Time

	PublicType PublicType `gorm:"foreignKey:PublicTypeID"`
	Container  Container  `gorm:"constraint:OnDelete:CASCADE;foreignKey:ContainerID;references:ID"`
}

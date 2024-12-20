package model

import (
	"time"
)

type Container struct {
	ID         uint   `gorm:"primary_key"`
	Name       string `gorm:"unique"`
	Body       string `gorm:"type:text"`
	IsActive   bool   `gorm:"default:true"`
	ComposerID uint   `gorm:"index;default:1"` //Default 1 пока нет реализации разделения по composer
	CreatedAt  time.Time
	UpdatedAt  time.Time

	PublicData []PublicData `gorm:"foreignKey:ContainerID;constraint:OnDelete:CASCADE"`
}

func (c *Container) GetPublicDataByName(name string) (PublicData, bool) {
	for _, data := range c.PublicData {
		if data.PublicType.Name == name {
			return data, true
		}
	}

	return PublicData{}, false
}

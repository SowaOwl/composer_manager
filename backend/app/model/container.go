package model

import (
	"composer_vue/backend/util/debug"
	"time"
)

type Container struct {
	ID         uint   `gorm:"primary_key"`
	Name       string `gorm:"unique"`
	Body       string `gorm:"type:text"`
	ComposerID uint   `gorm:"index"`
	CreatedAt  time.Time
	UpdatedAt  time.Time

	PublicData []PublicData `gorm:"foreignKey:ContainerID"`
}

func (c *Container) GetPublicDataByName(name string) (PublicData, bool) {
	for _, data := range c.PublicData {
		if data.PublicType.Name == name {
			debug.JsonPrint(data)
			debug.JsonPrint(name)
			return data, true
		}
	}

	return PublicData{}, false
}

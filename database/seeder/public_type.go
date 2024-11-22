package seeder

import (
	"composer_vue/backend/app/model"
	"gorm.io/gorm"
)

func publicTypesSeed(db *gorm.DB) {
	publicTypes := []model.PublicType{
		{
			ID:         1,
			Name:       "port",
			ManyName:   "ports",
			IsTabulate: true,
		},
		{
			ID:         2,
			Name:       "volume",
			ManyName:   "volumes",
			IsTabulate: true,
		},
		{
			ID:         3,
			Name:       "working_dir",
			ManyName:   "",
			IsTabulate: false,
		},
	}

	for _, publicType := range publicTypes {
		if err := db.FirstOrCreate(&publicType, model.PublicType{ID: publicType.ID}).Error; err != nil {
			panic(err)
		}
	}
}

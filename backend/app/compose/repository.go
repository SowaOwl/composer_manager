package compose

import (
	"composer_vue/backend/app/model"
	"gorm.io/gorm"
)

type Repository interface {
	GetTypes() []model.PublicType
	GetContainers() []model.Container
	GetNetworks() []model.Network
}

type GormRepository struct {
	db *gorm.DB
}

func NewGormComposeRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{
		db: db,
	}
}

func (r *GormRepository) GetTypes() []model.PublicType {
	var types []model.PublicType
	r.db.Find(&types)
	return types
}

func (r *GormRepository) GetContainers() []model.Container {
	var containers []model.Container
	r.db.Preload("PublicData.PublicType").Find(&containers)
	return containers
}

func (r *GormRepository) GetNetworks() []model.Network {
	var networks []model.Network
	r.db.Find(&networks)
	return networks
}

package compose

import (
	"composer_vue/backend/app/model"
	"gorm.io/gorm"
)

type Repository interface {
	GetTypes() []model.PublicType
	GetContainers() []model.Container
	GetNetworks() []model.Network
	GetContainerByID(id int64) *model.Container
	GetNetworkByID(id int64) *model.Network
	GetTypeByID(id int64) *model.PublicType
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

func (r *GormRepository) GetContainerByID(id int64) *model.Container {
	var container model.Container
	r.db.Preload("PublicData.PublicType").First(&container, id)
	return &container
}

func (r *GormRepository) GetNetworkByID(id int64) *model.Network {
	var network model.Network
	r.db.First(&network, id)
	return &network
}

func (r *GormRepository) GetTypeByID(id int64) *model.PublicType {
	var cType model.PublicType
	r.db.First(&cType, id)
	return &cType
}

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

	UpdateContainer(*model.Container) error

	CreateContainer(*model.Container) (*model.Container, error)
	CreatePublicData(*model.PublicData) (*model.PublicData, error)
	CreatePublicType(*model.PublicType) error

	DeleteContainerByID(id int64) error
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

func (r *GormRepository) UpdateContainer(c *model.Container) error {
	return r.db.Save(c).Error
}

func (r *GormRepository) CreateContainer(c *model.Container) (*model.Container, error) {
	if err := r.db.FirstOrCreate(&c, model.Container{Name: c.Name}).Error; err != nil {
		return nil, err
	}
	return c, nil
}

func (r *GormRepository) CreatePublicData(c *model.PublicData) (*model.PublicData, error) {
	if err := r.db.Create(&c).Error; err != nil {
		return nil, err
	}
	return c, nil
}

func (r *GormRepository) CreatePublicType(c *model.PublicType) error {
	return r.db.Create(&c).Error
}

func (r *GormRepository) DeleteContainerByID(id int64) error {
	var container model.Container
	r.db.First(&container, id)
	return r.db.Delete(&container).Error
}

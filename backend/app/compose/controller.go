package compose

import "composer_vue/backend/app/model"

type Controller interface {
	GetAllTypes() []model.PublicType
	GetAllContainers() []model.Container
	GetContainerByID(id int) *model.Container
	GetTypeByID(id int) *model.PublicType
	GetNetworkByID(id int) *model.Network
	UpdateContainer(request UpdateContainerRequest) error
	CreateContainer(request CreateContainerRequest) error
	DeleteContainer(id int) error
	GenerateDockerCompose() error
}

type ControllerImpl struct {
	service Service
	repo    Repository
}

func NewController(service Service, repo Repository) *ControllerImpl {
	return &ControllerImpl{
		service: service,
		repo:    repo,
	}
}

func (c *ControllerImpl) GetAllTypes() []model.PublicType {
	return c.repo.GetTypes()
}

func (c *ControllerImpl) GetAllContainers() []model.Container {
	return c.repo.GetContainers()
}

func (c *ControllerImpl) GetContainerByID(id int64) *model.Container {
	return c.repo.GetContainerByID(id)
}

func (c *ControllerImpl) GetTypeByID(id int64) *model.PublicType {
	return c.repo.GetTypeByID(id)
}

func (c *ControllerImpl) GetNetworkByID(id int64) *model.Network {
	return c.repo.GetNetworkByID(id)
}

func (c *ControllerImpl) CreateContainer(request CreateContainerRequest) error {
	return c.service.CreateContainer(request)
}

func (c *ControllerImpl) GenerateDockerCompose() error {
	return c.service.GenerateDockerCompose()
}

func (c *ControllerImpl) DeleteContainer(id int) error {
	return c.repo.DeleteContainerByID(int64(id))
}

package compose

import (
	"composer_vue/backend/app/model"
)

type Service interface {
	GenerateDockerCompose() error
	CreateContainer(request CreateContainerRequest) error
}

type ComposeService struct {
	repo Repository
}

func NewGormCompose(repo Repository) *ComposeService {
	return &ComposeService{
		repo: repo,
	}
}

func (c ComposeService) GenerateDockerCompose() error {
	containers := c.repo.GetContainers()
	types := c.repo.GetTypes()
	networks := c.repo.GetNetworks()

	err := generateComposeFile(containers, types, networks)
	if err != nil {
		return err
	}

	return nil
}

func (c ComposeService) CreateContainer(request CreateContainerRequest) error {
	container := model.Container{
		Name: request.Name,
		Body: request.Body,
	}

	result, err := c.repo.CreateContainer(&container)
	if err != nil {
		return err
	}

	for _, data := range request.PublicTypes {
		_, err := c.repo.CreatePublicData(&model.PublicData{
			ContainerID:  result.ID,
			PublicTypeID: data.TypeID,
			Data:         data.Data,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

package compose

type Service interface {
	GenerateDockerCompose() error
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

package compose

import (
	"composer_vue/backend/app/model"
	"composer_vue/backend/util/debug"
	"os"
	"strings"
)

type Service interface {
	GetAllTypes() []model.PublicType
	GetAllContainers() []model.Container
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

func (c ComposeService) GetAllTypes() []model.PublicType {
	return c.repo.GetTypes()
}

func (c ComposeService) GetAllContainers() []model.Container {
	return c.repo.GetContainers()
}

func (c ComposeService) GenerateDockerCompose() error {
	file, err := os.Create("docker-compose.yml")
	if err != nil {
		return err
	}

	_, err = file.WriteString("version: '3'\n")
	if err != nil {
		return err
	}

	containers := c.repo.GetContainers()
	types := c.repo.GetTypes()
	err = c.writeServices(file, containers, types)

	networks := c.repo.GetNetworks()
	err = c.writeNetworks(file, networks)

	return nil
}

func (c ComposeService) writeServices(file *os.File, containers []model.Container, types []model.PublicType) error {
	_, err := file.WriteString("services:\n")
	if err != nil {
		return err
	}

	for _, container := range containers {
		replacedText, err := c.replaceText(container.Body, types, container)
		if err != nil {
			return err
		}

		_, err = file.WriteString(replacedText + "\n\n")
		if err != nil {
			return err
		}
	}
	return nil
}

func (c ComposeService) writeNetworks(file *os.File, networks []model.Network) error {
	_, err := file.WriteString("networks:\n")
	if err != nil {
		return err
	}

	for _, network := range networks {
		_, err := file.WriteString(network.Body + "\n\n")
		if err != nil {
			return err
		}
	}

	return nil
}

func (c ComposeService) replaceText(text string, types []model.PublicType, container model.Container) (string, error) {
	replacedText := text

	for _, containerType := range types {
		find := "{" + containerType.Name + "}"
		if strings.Contains(replacedText, find) {
			if data, ok := container.GetPublicDataByName(containerType.Name); ok {
				debug.JsonPrint(data)
				textToReplace := ""
				if containerType.IsTabulate {
					textToReplace += "\n"
				}
				textToReplace += c.getTextDecoratorForBridge(containerType, data.Data)
				replacedText = strings.ReplaceAll(replacedText, find, textToReplace)
			}
		}

		find = "{" + containerType.ManyName + "}"
		if strings.Contains(replacedText, find) {
			generatedText := c.generateText(containerType)
			replacedText = strings.ReplaceAll(replacedText, find, generatedText)
		}
	}

	return replacedText, nil
}

func (c ComposeService) getTextDecoratorForBridge(cType model.PublicType, text string) string {
	switch cType.Name {
	case "port":
		return "\t\t\t - \"" + text + "\""
	case "volume":
		return "\t\t\t - " + text + ""
	default:
		return text
	}
}

func (c ComposeService) generateText(cType model.PublicType) string {
	generatedText := ""

	containers := c.repo.GetContainers()
	for _, container := range containers {
		if data, ok := container.GetPublicDataByName(cType.Name); ok {
			generatedText += "\n"
			generatedText += c.getTextDecoratorForBridge(cType, data.Data)
		}
	}

	return generatedText
}

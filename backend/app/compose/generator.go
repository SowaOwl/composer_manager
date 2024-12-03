package compose

import (
	"composer_vue/backend/app/model"
	"os"
	"strings"
)

func generateComposeFile(containers []model.Container, types []model.PublicType, networks []model.Network) error {
	file, err := os.Create("docker-compose.yml")
	if err != nil {
		return err
	}

	_, err = file.WriteString("version: '3'\n")
	if err != nil {
		return err
	}

	err = writeServices(file, containers, types)
	if err != nil {
		return err
	}

	err = writeNetworks(file, networks)
	if err != nil {
		return err
	}

	return nil
}

func writeServices(file *os.File, containers []model.Container, types []model.PublicType) error {
	_, err := file.WriteString("services:\n")
	if err != nil {
		return err
	}

	for _, container := range containers {
		if !container.IsActive {
			continue
		}

		replacedText, err := replaceText(container.Body, types, container, containers)
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

func writeNetworks(file *os.File, networks []model.Network) error {
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

func replaceText(text string, types []model.PublicType, container model.Container, allContainers []model.Container) (string, error) {
	replacedText := text

	for _, containerType := range types {
		find := "{" + containerType.Name + "}"
		if strings.Contains(replacedText, find) {
			if data, ok := container.GetPublicDataByName(containerType.Name); ok {
				textToReplace := ""
				if containerType.IsTabulate {
					textToReplace += "\n"
				}
				textToReplace += getTextDecoratorForBridge(containerType, data.Data)
				replacedText = strings.ReplaceAll(replacedText, find, textToReplace)
			}
		}

		find = "{" + containerType.ManyName + "}"
		if strings.Contains(replacedText, find) {
			generatedText := generateText(containerType, allContainers)
			replacedText = strings.ReplaceAll(replacedText, find, generatedText)
		}
	}

	return replacedText, nil
}

func getTextDecoratorForBridge(cType model.PublicType, text string) string {
	switch cType.Name {
	case "port":
		return "\t\t\t - \"" + text + "\""
	case "volume":
		return "\t\t\t - " + text + ""
	default:
		return text
	}
}

func generateText(cType model.PublicType, containers []model.Container) string {
	generatedText := ""

	for _, container := range containers {
		if !container.IsActive {
			continue
		}
		if data, ok := container.GetPublicDataByName(cType.Name); ok {
			generatedText += "\n"
			generatedText += getTextDecoratorForBridge(cType, data.Data)
		}
	}

	return generatedText
}

package seeder

import (
	"composer_vue/backend/app/model"
	"gorm.io/gorm"
)

func testSeed(db *gorm.DB) {
	composer := model.Composer{
		ID:   1,
		Path: "test/test",
	}

	containers := []model.Container{
		{
			ID:         1,
			Name:       "main",
			Body:       "\twebserver:\n\t\timage: nginx:alpine\n\t\trestart: unless-stopped\n\t\ttty: true\n\t\tcontainer_name: webserver\n\t\tports: {ports}\n\t\tvolumes: {volumes}\n\t\tnetworks: \n\t\t\t- dev.a-i.kz",
			ComposerID: composer.ID,
		},
		{
			ID:         2,
			Name:       "test",
			Body:       "\ttest:\n\t\tbuild:\n\t\t\tcontext: apps/test\n\t\t\tdockerfile: Dockerfile\n\t\trestart: unless-stopped\n\t\ttty: true\n\t\tworking_dir: test/test\n\t\tvolumes: {volume}\n\t\tnetworks: \n\t\t\t- dev.a-i.kz",
			ComposerID: composer.ID,
		},
		{
			ID:         3,
			Name:       "test2",
			Body:       "\tnew_test:\n\t\tbuild:\n\t\t\tcontext: apps/test\n\t\t\tdockerfile: Dockerfile\n\t\trestart: unless-stopped\n\t\ttty: true\n\t\tworking_dir: test/test\n\t\tvolumes: {volume}\n\t\tports: {port}\n\t\tnetworks: \n\t\t\t- dev.a-i.kz",
			ComposerID: composer.ID,
		},
	}

	publicData := []model.PublicData{
		{
			PublicTypeID: 1,
			ContainerID:  containers[1].ID,
			Data:         "2022:2022",
		},
		{
			PublicTypeID: 2,
			ContainerID:  containers[1].ID,
			Data:         "test/test/:/var/test",
		},
		{
			PublicTypeID: 1,
			ContainerID:  containers[2].ID,
			Data:         "3333:3333",
		},
		{
			PublicTypeID: 2,
			ContainerID:  containers[2].ID,
			Data:         "test/test2/:/var/test",
		},
	}

	networks := []model.Network{
		{
			Name:       "dev.a-i.kz",
			Body:       "\tdev.a-i.kz:\n\t\tdriver: bridge\n\t\tipam:\n\t\t\tconfig:\n\t\t\t\t- subnet: \"172.18.0.0/24\"",
			ComposerID: composer.ID,
		},
	}

	err := db.FirstOrCreate(&composer, model.Composer{ID: composer.ID}).Error
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		err := db.FirstOrCreate(&container, model.Container{ID: container.ID}).Error
		if err != nil {
			panic(err)
		}
	}

	for _, publicData := range publicData {
		err := db.FirstOrCreate(&publicData, model.PublicData{ContainerID: publicData.ContainerID, Data: publicData.Data}).Error
		if err != nil {
			panic(err)
		}
	}

	for _, network := range networks {
		err := db.FirstOrCreate(&network, model.Network{Name: network.Name}).Error
		if err != nil {
			panic(err)
		}
	}
}

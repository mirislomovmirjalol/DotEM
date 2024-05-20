package service

import (
	"github/mirislomovmirjalol/DotEM/internal/store"
	"io/ioutil"
)

func AddProject(projectName string, filePath string) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	boltDBStore := store.GetStore()
	defer boltDBStore.Close()

	err = boltDBStore.Set("Projects", projectName, data)
	if err != nil {
		return err
	}
	return nil
}

func GetProject(projectName string) ([]byte, error) {
	boltDBStore := store.GetStore()
	defer boltDBStore.Close()

	data, err := boltDBStore.Get("Projects", projectName)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func GetAllProjects() ([]string, error) {
	boltDBStore := store.GetStore()
	defer boltDBStore.Close()

	projects, err := boltDBStore.GetKeys("Projects")
	if err != nil {
		return nil, err
	}

	return projects, nil
}

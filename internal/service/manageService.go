package service

import (
	"encoding/json"
	"github/mirislomovmirjalol/DotEM/internal/store"
	"os"
)

func WriteDataToFile(data []byte, filePath string) error {
	err := os.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func ExportData(filePath string) error {
	boltDBStore := store.GetStore()
	defer boltDBStore.Close()
	data, err := boltDBStore.GetKeysWithValue("Projects")
	if err != nil {
		return err
	}
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, jsonData, 0644)
}

func ImportData(filePath string) error {
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var envData map[string]string
	err = json.Unmarshal(fileData, &envData)
	if err != nil {
		return err
	}
	boltDBStore := store.GetStore()
	defer boltDBStore.Close()

	for key, value := range envData {
		err = boltDBStore.Set("Projects", key, []byte(value))
		if err != nil {
			return err
		}
	}

	return nil
}

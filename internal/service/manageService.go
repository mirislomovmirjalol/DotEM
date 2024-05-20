package service

import "os"

func exportData() ([]byte, error) {
	return nil, nil
}

func WriteDataToFile(data []byte, filePath string) error {
	err := os.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

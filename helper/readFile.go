package helper

import (
	"encoding/csv"
	"os"
)

func ReadFile(filePath string) ([][]string, error) {

	file, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	text, err := csv.NewReader(file).ReadAll()

	if err != nil {
		return nil, err
	}

	return text, nil

}

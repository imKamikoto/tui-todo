package main

import (
	"encoding/json"
	"errors"
	"os"
)

type Storage[T any] struct {
	fileName string
}

func NewStorage[T any](fileName string) (*Storage[T], error) {

	if fileName == "" {
		return nil, errors.New("Invalid file name")
	}
	return &Storage[T]{fileName: fileName}, nil
}

func (storage *Storage[T]) Save(data *T) error {

	fileData, error := json.Marshal(data)

	if error != nil {
		return error
	}

	return os.WriteFile(storage.fileName, fileData, 0644)
}

func (storage *Storage[T]) Load(data *T) error {

	fileData, error := os.ReadFile(storage.fileName)

	if error != nil {
		return error
	}

	return json.Unmarshal(fileData, data)
}

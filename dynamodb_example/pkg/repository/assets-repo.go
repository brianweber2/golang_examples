package repository

import "github.com/brianweber2/golang_examples/dynamodb_example/pkg/models"

type AssetsRepository interface {
	FindAll() ([]models.Asset, error)
	FindById(id string) (*models.Asset, error)
}

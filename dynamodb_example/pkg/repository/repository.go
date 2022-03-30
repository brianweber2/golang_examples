package repository

import "github.com/brianweber2/golang_examples/dynamodb_example/pkg/models"

type OpexRepository interface {
	FindAllAssets() ([]models.Asset, error)
	FindAssetById(id string) (*models.Asset, error)
	FindAllAssetsMetrics() ([]models.AssetMetrics, error)
	FindAssetMetricsById(id string) (*models.AssetMetrics, error)
}

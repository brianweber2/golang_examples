package main

import (
	"fmt"

	"github.com/brianweber2/golang_examples/dynamodb_example/pkg/repository"
)

var (
	repo repository.OpexRepository = repository.NewDynamoDBRepository()
)

func main() {
	assets, err := repo.FindAllAssets()
	if err != nil {
		panic(err)
	}
	fmt.Println("Number of assets returned: ", len(assets))

	assetId := "8937411052819864529"

	asset, err := repo.FindAssetById(assetId)
	if err != nil {
		panic(err)
	}
	fmt.Println("Data for AssetID", assetId, ":", asset)

	assetsMetrics, err := repo.FindAllAssetsMetrics()
	if err != nil {
		panic(err)
	}
	fmt.Println("Number of assets with metrics returned: ", len(assetsMetrics))

	assetMetrics, err := repo.FindAssetMetricsById(assetId)
	if err != nil {
		panic(err)
	}
	fmt.Println("Metrics data for AssetID", assetId, ":", assetMetrics)
}

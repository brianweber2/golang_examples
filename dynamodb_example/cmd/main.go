package main

import (
	"context"
	"fmt"

	internal "github.com/brianweber2/golang_examples/dynamodb_example/internal/db"
	dbConfig "github.com/brianweber2/golang_examples/dynamodb_example/internal/db/config"
	"github.com/brianweber2/golang_examples/dynamodb_example/pkg/repository"
)

var (
	ctx   context.Context           = context.Background()
	dbCfg *dbConfig.DynamoDBConfig  = dbConfig.NewDbConfig(ctx)
	repo  repository.OpexRepository = repository.NewDynamoDBRepository(internal.CreateDynamoDBClient(ctx, dbCfg))
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

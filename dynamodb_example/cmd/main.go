package main

import (
	"fmt"

	"github.com/brianweber2/golang_examples/dynamodb_example/pkg/repository"
)

var (
	assetRepository repository.AssetsRepository = repository.NewDynamoDBRepository()
)

func main() {
	assets, err := assetRepository.FindAll()
	if err != nil {
		panic(err)
	}

	fmt.Println("Number of assets returned: ", len(assets))

	assetId := "8937411052819864529"
	asset, err := assetRepository.FindById(assetId)
	if err != nil {
		panic(err)
	}

	fmt.Println("Data for AssetID", assetId, ":", asset)
}

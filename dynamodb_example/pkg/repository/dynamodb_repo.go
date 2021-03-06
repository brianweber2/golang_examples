package repository

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
	dynamoDB "github.com/brianweber2/golang_examples/dynamodb_example/internal/db"
	"github.com/brianweber2/golang_examples/dynamodb_example/pkg/models"
)

const (
	MetricsTable      = "S360-METRICS"
	S360ServicesTable = "S360-SERVICES"
)

// Tables - DynamoDB Table Names
type Tables struct {
	MetricTable       string
	S360ServicesTable string
}

// tables instantiates Tables struct with the env variables
var tables = Tables{
	MetricTable:       os.Getenv(MetricsTable),
	S360ServicesTable: os.Getenv(S360ServicesTable),
}

type dynamoDBRepo struct {
	tableNames  *Tables
	DynamoDbApi dynamoDB.DynamoDbApi
}

// NewDynamoDBRepository is the constructor function for the repo
func NewDynamoDBRepository(client dynamoDB.DynamoDbApi) OpexRepository {
	// ctx := context.Background()
	// dynamoDBConfig := dynamoDBConfig.NewDbConfig(ctx)

	return &dynamoDBRepo{
		tableNames:  &tables,
		DynamoDbApi: client,
		// DynamoDbApi: dynamoDB.CreateDynamoDBClient(ctx, dynamoDBConfig),

	}
}

func (repo *dynamoDBRepo) FindAllAssets() ([]models.Asset, error) {
	// Build the query input parameters
	params := dynamodb.ScanInput{
		TableName: aws.String(repo.tableNames.S360ServicesTable),
	}

	// Make the DynamoDB Query API call
	results, err := repo.DynamoDbApi.Scan(context.Background(), &params)
	if err != nil {
		return nil, err
	}

	// Create a assets array and add all the existing assets
	var assets []models.Asset = []models.Asset{}
	for _, i := range results.Items {
		asset := models.Asset{}

		err = attributevalue.UnmarshalMap(i, &asset)
		if err != nil {
			panic(err)
		}
		assets = append(assets, asset)
	}

	// Return the assets array
	return assets, nil
}

func (repo *dynamoDBRepo) FindAssetById(id string) (*models.Asset, error) {
	// Build the query input parameters
	params := dynamodb.GetItemInput{
		Key: map[string]types.AttributeValue{
			"AssetID": &types.AttributeValueMemberS{Value: id},
		},
		TableName: aws.String(repo.tableNames.S360ServicesTable),
	}

	// Get the item by ID
	result, err := repo.DynamoDbApi.GetItem(context.Background(), &params)
	if err != nil {
		return nil, err
	}

	// Map the dynamodb element to the asset structure
	asset := models.Asset{}

	err = attributevalue.UnmarshalMap(result.Item, &asset)
	if err != nil {
		panic(err)
	}

	return &asset, nil
}

func (repo *dynamoDBRepo) FindAllAssetsMetrics() ([]models.AssetMetrics, error) {
	// Build the query input parameters
	params := dynamodb.ScanInput{
		TableName: aws.String(repo.tableNames.MetricTable),
	}

	// Make the DynamoDB Query API call
	results, err := repo.DynamoDbApi.Scan(context.Background(), &params)
	if err != nil {
		return nil, err
	}

	// Create a assetsMetrics array and add all the existing assetMetrics
	var assetsMetrics []models.AssetMetrics = []models.AssetMetrics{}
	for _, i := range results.Items {
		assetMetrics := models.AssetMetrics{}

		err = attributevalue.UnmarshalMap(i, &assetMetrics)
		if err != nil {
			panic(err)
		}
		assetsMetrics = append(assetsMetrics, assetMetrics)
	}

	// Return the assets array
	return assetsMetrics, nil
}

func (repo *dynamoDBRepo) FindAssetMetricsById(id string) (*models.AssetMetrics, error) {
	// Build the query input parameters
	params := dynamodb.GetItemInput{
		Key: map[string]types.AttributeValue{
			"AssetID": &types.AttributeValueMemberS{Value: id},
		},
		TableName: aws.String(repo.tableNames.MetricTable),
	}

	// Get the item by ID
	result, err := repo.DynamoDbApi.GetItem(context.Background(), &params)
	if err != nil {
		return nil, err
	}

	// Map the dynamodb element to the asset metrics structure
	assetMetrics := models.AssetMetrics{}

	err = attributevalue.UnmarshalMap(result.Item, &assetMetrics)
	if err != nil {
		panic(err)
	}

	return &assetMetrics, nil
}

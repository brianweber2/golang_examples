package repository

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/brianweber2/golang_examples/dynamodb_example/pkg/models"
)

type dynamoDBRepo struct {
	tableName string
}

// NewDynamoDBRepository is the constructor function for the repo
func NewDynamoDBRepository() AssetsRepository {
	return &dynamoDBRepo{
		tableName: "s360-services-preprod",
	}
}

func createDynamoDBClient() *dynamodb.Client {
	cfg, _ := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("cgsre-ppd-aws"), config.WithRegion("us-west-2"))
	return dynamodb.NewFromConfig(cfg)
}

func (repo *dynamoDBRepo) FindAll() ([]models.Asset, error) {
	// Get a new DynamoDB client
	dynamoDBClient := createDynamoDBClient()

	// Build the query input parameters
	params := dynamodb.ScanInput{
		TableName: aws.String(repo.tableName),
	}

	// Make the DynamoDB Query API call
	results, err := dynamoDBClient.Scan(context.Background(), &params)
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

func (repo *dynamoDBRepo) FindById(id string) (*models.Asset, error) {
	// Get a new DynamoDB client
	dynamoDBClient := createDynamoDBClient()

	// Build the query input parameters
	params := dynamodb.GetItemInput{
		Key: map[string]types.AttributeValue{
			"AssetID": &types.AttributeValueMemberS{Value: id},
		},
		TableName: aws.String(repo.tableName),
	}

	// Get the item by ID
	result, err := dynamoDBClient.GetItem(context.Background(), &params)
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

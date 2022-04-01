package internal

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	dynamodbCfg "github.com/brianweber2/golang_examples/dynamodb_example/internal/db/config"
)

type DynamoDbApi interface {
	GetItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)
	Scan(ctx context.Context, params *dynamodb.ScanInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ScanOutput, error)
}

func CreateDynamoDBClient(ctx context.Context, dynamoDBConfig *dynamodbCfg.DynamoDBConfig) DynamoDbApi {
	cfg, _ := config.LoadDefaultConfig(ctx, config.WithSharedConfigProfile("cgsre-ppd-aws"), config.WithRegion("us-west-2"))
	return dynamodb.NewFromConfig(cfg)
}

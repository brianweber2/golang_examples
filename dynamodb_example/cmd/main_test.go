package main

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	internal "github.com/brianweber2/golang_examples/dynamodb_example/internal/db"
)

type dynamoDbClientMock struct {
	internal.DynamoDbApi
}

func (c dynamoDbClientMock) Scan(ctx context.Context, params *dynamodb.ScanInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ScanOutput, error) {
	return nil, nil
}

func TestMain(t *testing.T) {
}

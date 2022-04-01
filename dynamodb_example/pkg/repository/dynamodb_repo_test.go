package repository

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/stretchr/testify/assert"
)

type MockDynamoDbApi struct{}

func createTestAssets() []map[string]types.AttributeValue {
	items := make([]map[string]types.AttributeValue, 0)

	for i := 1; i <= 25; i++ {
		item := map[string]types.AttributeValue{
			"AssetID":     &types.AttributeValueMemberS{Value: strconv.Itoa(i)},
			"ServiceName": &types.AttributeValueMemberS{Value: fmt.Sprintf("ServiceName%d", i)},
			// "Metrics":     &types.AttributeValueMemberL{Value: []types.AttributeValue{}},
		}
		items = append(items, item)
	}
	return items
}

func (m *MockDynamoDbApi) Scan(ctx context.Context, params *dynamodb.ScanInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ScanOutput, error) {
	output := dynamodb.ScanOutput{
		Items: createTestAssets(),
	}
	return &output, nil
}

func (m *MockDynamoDbApi) GetItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error) {
	output := dynamodb.GetItemOutput{
		Item: createTestAssets()[0],
	}
	return &output, nil
}

var mockRepo OpexRepository = NewDynamoDBRepository(&MockDynamoDbApi{})

func TestFindAllAssets(t *testing.T) {
	// Need to mock repo.DynamoDbApi.Scan
	assets, err := mockRepo.FindAllAssets()
	if err != nil {
		t.Errorf("expected no error when getting all assets, but got: %v", err)
	}

	assert.Equal(t, 25, len(assets))
}

func TestFindAssetById(t *testing.T) {
	// Need to mock repo.DynamoDbApi.GetItem
	asset, err := mockRepo.FindAssetById("1")
	if err != nil {
		t.Errorf("expected no error when getting asset, but got: %v", err)
	}

	assert.Equal(t, "1", asset.AssetID)
}

func TestFindAllAssetsMetrics(t *testing.T) {
	// Need to mock repo.DynamoDbApi.Scan
	assetsMetrics, err := mockRepo.FindAllAssetsMetrics()
	if err != nil {
		t.Errorf("expected no error when getting all assets metrics, but got: %v", err)
	}

	assert.Equal(t, 25, len(assetsMetrics))
}

func TestFindAssetMetricsById(t *testing.T) {
	// Need to mock repo.DynamoDbApi.GetItem
	assetMetrics, err := mockRepo.FindAssetMetricsById("1")
	if err != nil {
		t.Errorf("expected no error when getting asset metrics, but got: %v", err)
	}

	assert.Equal(t, "1", assetMetrics.AssetID)
}

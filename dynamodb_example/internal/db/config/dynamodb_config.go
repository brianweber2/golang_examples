package internal

import (
	"context"
	"os"
	"time"
)

type DynamoDBConfig struct {
	ConnectToRemoteDB bool          `default:"false" split_words:"true"`
	LocalDBEndpoint   string        `split_words:"true"`
	IAMRole           string        `split_words:"true"`
	Region            string        `default:"us-west-2"`
	Environment       string        `default:"dev"`
	QueryTimeout      time.Duration `default:"0.5s" split_words:"true"`
	MaxBackOff        time.Duration `default:"50ms" split_words:"true"`
	MaxAttempts       int           `default:"4" split_words:"true"`
	TimeToLive        int64         `default:"3600" split_words:"true"`
	BatchSize         int           `default:"25" split_words:"true"`
	AWSProfile        string        `default:"25" split_words:"true"`
}

func NewDbConfig(ctx context.Context) *DynamoDBConfig {
	dbConfig := DynamoDBConfig{
		QueryTimeout:      500 * time.Millisecond,
		MaxBackOff:        50 * time.Millisecond,
		MaxAttempts:       2,
		TimeToLive:        int64(3600),
		BatchSize:         25,
		ConnectToRemoteDB: false,
	}

	assumeRoleArn := os.Getenv("AWS_IAM_ROLE_ARN")
	awsProfile := os.Getenv("AWS_PROFILE_DYNAMO")
	awsLocalDynamo := os.Getenv("AWS_LOCAL_DYNAMO")
	region := os.Getenv("AWS_REGION")

	if region != "" {
		dbConfig.Region = region
	}

	if assumeRoleArn != "" {
		dbConfig.IAMRole = assumeRoleArn
	} else if awsProfile != "" {
		dbConfig.AWSProfile = awsProfile
		dbConfig.ConnectToRemoteDB = true
	} else {
		dbConfig.LocalDBEndpoint = awsLocalDynamo
		dbConfig.ConnectToRemoteDB = false
	}

	return &dbConfig
}

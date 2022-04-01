package internal

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDbConfigWithRegionAndIamRole(t *testing.T) {
	os.Setenv("AWS_REGION", "us-east-2")
	os.Setenv("AWS_IAM_ROLE_ARN", "arn")
	defer os.Unsetenv("AWS_REGION")
	defer os.Unsetenv("AWS_IAM_ROLE_ARN")

	dynamodbConfig := NewDbConfig(context.TODO())

	assert.Equal(t, "us-east-2", dynamodbConfig.Region)
	assert.Equal(t, "arn", dynamodbConfig.IAMRole)
}

func TestNewDbConfigWithAwsProfile(t *testing.T) {
	os.Setenv("AWS_PROFILE_DYNAMO", "profile")
	defer os.Unsetenv("AWS_PROFILE_DYNAMO")

	dynamodbConfig := NewDbConfig(context.TODO())

	assert.Equal(t, "profile", dynamodbConfig.AWSProfile)
}

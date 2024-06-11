// Copyright (c) The OpenTofu Authors
// SPDX-License-Identifier: MPL-2.0
// Copyright (c) 2023 HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package testutils_test

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/opentofu/opentofu/internal/testutils"
)

func TestDynamoDBService(t *testing.T) {
	ctx := testutils.Context(t)
	var dynamoDBService testutils.AWSDynamoDBTestService = testutils.AWS(t)
	var dynamoDBClient = dynamodb.NewFromConfig(dynamoDBService.ConfigV2())
	t.Logf("🔍 Checking if the DynamoDB table from the AWS test service can be described...")
	if _, err := dynamoDBClient.DescribeTable(ctx, &dynamodb.DescribeTableInput{
		TableName: aws.String(dynamoDBService.DynamoDBTable()),
	}); err != nil {
		t.Fatalf("❌ %v", err)
	}
	t.Logf("✅ DynamoDB works as intended.")
}

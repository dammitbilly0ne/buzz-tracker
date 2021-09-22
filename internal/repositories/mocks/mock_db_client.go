package mocks

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/stretchr/testify/mock"
)

type MockDynamoDBClient struct {
	dynamodbiface.DynamoDBAPI
	mock.Mock
}

func (m *MockDynamoDBClient) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	args := m.Called(input)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*dynamodb.PutItemOutput), nil
}

func (m *MockDynamoDBClient) GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	args := m.Called(input)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*dynamodb.GetItemOutput), nil
}

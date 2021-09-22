package breweries

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/dammitbilly0ne/buzz-tracker/internal/entities"
	"testing"

	"github.com/dammitbilly0ne/buzz-tracker/internal/repositories/mocks"

	"github.com/stretchr/testify/assert"
)

const (
	testTable = "testTable"
	idField   = "id"
)

func setupFixture() *Dynamo {
	return &Dynamo{
		client:    &mocks.MockDynamoDBClient{},
		tableName: testTable,
	}
}

func Test_NewDynamo(t *testing.T) {
	t.Run("it returns a valid brewery dynamo", func(t *testing.T) {
		actual, err := NewDynamo(&DynamoConfig{
			Client:    &mocks.MockDynamoDBClient{},
			TableName: testTable,
		})

		assert.Nil(t, err)
		assert.NotNil(t, actual)
		assert.NotNil(t, actual.client)
		assert.Equal(t, testTable, actual.tableName)
	})

	t.Run("it requires a brewery config", func(t *testing.T) {
		actual, err := NewDynamo(nil)

		expErr := errors.New("cfg is required")
		assert.Nil(t, actual)
		assert.NotNil(t, err)
		assert.Equal(t, expErr, err)
	})

	t.Run("it requires a Client", func(t *testing.T) {
		actual, err := NewDynamo(&DynamoConfig{})

		expErr := errors.New("Client is required")
		assert.Nil(t, actual)
		assert.NotNil(t, err)
		assert.Equal(t, expErr, err)
	})

	t.Run("it requires a TableName", func(t *testing.T) {
		actual, err := NewDynamo(&DynamoConfig{
			Client: &mocks.MockDynamoDBClient{},
		})

		expErr := errors.New("TableName is required")
		assert.Nil(t, actual)
		assert.NotNil(t, err)
		assert.Equal(t, expErr, err)
	})
}

func TestDynamo_CreateBrewery(t *testing.T) {
	t.Run("brewery is required", func(t *testing.T) {
		repo := setupFixture()

		err := repo.CreateBrewery(nil)
		assert.Equal(t, errors.New("brewery is required"), err)
	})

	t.Run("brewery name is required", func(t *testing.T) {
		repo := setupFixture()

		err := repo.CreateBrewery(&entities.Brewery{})
		assert.Equal(t, errors.New("brewery name is required"), err)
	})

	t.Run("it creates a brewery", func(t *testing.T) {
		repo := setupFixture()
		m := repo.client.(*mocks.MockDynamoDBClient)

		testBrewery := &entities.Brewery{
			ID: "5",
			Name: "Denada Drinks",
			City: "Spokane",
			State: "Washington",
			Address: "123 Stacy Cr.",
			PhoneNumber: "(232)495869-4949",
			Website: "google.com",
		}

		mapped, _ := dynamodbattribute.MarshalMap(testBrewery)

		m.On("PutItem", &dynamodb.PutItemInput{
			TableName: aws.String(testTable),
			Item: mapped,
		}).Return(&dynamodb.PutItemOutput{}, nil)

		err := repo.CreateBrewery(testBrewery)

		assert.Nil(t, err)
		m.AssertExpectations(t)
	})

	t.Run("it returns an error when Dynamo errors", func(t *testing.T) {
		repo := setupFixture()
		m := repo.client.(*mocks.MockDynamoDBClient)

		testBrewery := &entities.Brewery{
			Name: "Denada Drinks",
		}

		mapped, _ := dynamodbattribute.MarshalMap(testBrewery)
		expErr := errors.New("dynamo has returned an error")
		m.On("PutItem", &dynamodb.PutItemInput{
			TableName: aws.String(testTable),
			Item: mapped,
		}).Return(nil, expErr)

		err := repo.CreateBrewery(testBrewery)

		assert.NotNil(t, err)
		assert.Equal(t, expErr, err)
		m.AssertExpectations(t)
	})
}

func TestDynamo_FindBrewery(t *testing.T) {
	t.Run("a brewery has been found", func(t *testing.T) {
		repo := setupFixture()
		m := repo.client.(*mocks.MockDynamoDBClient)
		input := &dynamodb.GetItemInput{
			TableName: &repo.tableName,
			Key: map[string]*dynamodb.AttributeValue{
				idField: {
					S: aws.String("Denada Drinks"),
				},
			},
		}

		expectedBrewery := &entities.Brewery{
			ID:"5",
		}

		mapped, _ := dynamodbattribute.MarshalMap(expectedBrewery)
		m.On("GetItem", input).Return(&dynamodb.GetItemOutput{
			Item: mapped,
		}, nil)

		actual, err := repo.FindBrewery("Denada Drinks")
		assert.Nil(t, err)
		assert.Equal(t, expectedBrewery, actual)
	})

	t.Run("it returns an error when the client errors", func(t *testing.T) {
		repo := setupFixture()
		m := repo.client.(*mocks.MockDynamoDBClient)

		testID := "Denada Drinks"

		input := &dynamodb.GetItemInput{
			TableName: aws.String(testTable),
			Key: map[string]*dynamodb.AttributeValue{
				"id": {
					S: aws.String(testID),
				},
			},

		}

		expErr := errors.New("dynamo down")

		m.On("GetItem", input).Return(nil, expErr)

		actual, err := repo.FindBrewery("Denada Drinks")

		assert.Nil(t, actual)
		assert.Equal(t, expErr, err)

		m.AssertExpectations(t)
	})
}
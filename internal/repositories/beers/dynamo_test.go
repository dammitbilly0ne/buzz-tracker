package beers

import (
	"errors"
	"testing"

	"github.com/dammitbilly0ne/buzz-tracker/internal/repositories/mocks"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/dammitbilly0ne/buzz-tracker/internal/entities"
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
	t.Run("it returns a valid dynamo", func(t *testing.T) {
		actual, err := NewDynamo(&DynamoConfig{
			Client:    &mocks.MockDynamoDBClient{},
			TableName: "testTable",
		})

		assert.Nil(t, err)
		assert.NotNil(t, actual)
		assert.NotNil(t, actual.client)
		assert.Equal(t, "testTable", actual.tableName)
	})

	t.Run("it requires a config", func(t *testing.T) {
		actual, err := NewDynamo(nil)

		expErr := errors.New("cfg is required")
		assert.Nil(t, actual)
		assert.NotNil(t, err)
		assert.Equal(t, expErr, err)
	})

	t.Run("it requires a Client", func(t *testing.T) {
		actual, err := NewDynamo(&DynamoConfig{})

		expERR := errors.New("cfg.Client is required")
		assert.Nil(t, actual)
		assert.NotNil(t, err)
		assert.Equal(t, expERR, err)
	})

	t.Run("it requires a TableName", func(t *testing.T) {
		actual, err := NewDynamo(&DynamoConfig{
			Client: &mocks.MockDynamoDBClient{},
		})

		expERR := errors.New("cfg.TableName is required")
		assert.Nil(t, actual)
		assert.NotNil(t, err)
		assert.Equal(t, expERR, err)
	})
}

func TestDynamo_StoreBeer(t *testing.T) {
	t.Run("beer is required", func(t *testing.T) {
		repo := setupFixture()

		err := repo.StoreBeer(nil)
		assert.Equal(t, errors.New(beerRequiredMsg), err)
	})

	t.Run("ID is required", func(t *testing.T) {
		repo := setupFixture()

		err := repo.StoreBeer(nil)
		assert.Equal(t, errors.New("ID is required"),err)
	})

	t.Run("beer name is required", func(t *testing.T) {
		repo := setupFixture()

		err := repo.StoreBeer(&entities.Beer{})
		assert.Equal(t, errors.New(beerNameRequiredMsg), err)
	})

	t.Run("it stores a beer", func(t *testing.T) {
		repo := setupFixture()
		m := repo.client.(*mocks.MockDynamoDBClient)

		testBeer := &entities.Beer{
			ID:      "3",
			Name:    "stan",
			Brewery: "spokane",
			Type:    "imperialIPA",
		}

		mapped, _ := dynamodbattribute.MarshalMap(testBeer)

		m.On("PutItem", &dynamodb.PutItemInput{
			TableName: aws.String("testTable"),
			Item:      mapped,
		}).Return(&dynamodb.PutItemOutput{}, nil)

		err := repo.StoreBeer(testBeer)

		assert.Nil(t, err)
		m.AssertExpectations(t)
	})

	t.Run("it returns an error when Dynamo errors", func(t *testing.T) {
		repo := setupFixture()
		m := repo.client.(*mocks.MockDynamoDBClient)

		testBeer := &entities.Beer{
			Name: "stan",
		}

		mapped, _ := dynamodbattribute.MarshalMap(testBeer)
		expectedErr := errors.New("dynammo has returned an error")
		m.On("PutItem", &dynamodb.PutItemInput{
			TableName: aws.String("testTable"),
			Item:      mapped,
		}).Return(nil, expectedErr)

		err := repo.StoreBeer(testBeer)

		assert.NotNil(t, err)
		assert.Equal(t, expectedErr, err)
		m.AssertExpectations(t)
	})
}

func Test_GetBeer(t *testing.T) {
	t.Run("a beer has been retrieved", func(t *testing.T) {
		repo := setupFixture()
		m := repo.client.(*mocks.MockDynamoDBClient)
		input := &dynamodb.GetItemInput{
			TableName: &repo.tableName,
			Key: map[string]*dynamodb.AttributeValue{
				idField: &dynamodb.AttributeValue{
					S: aws.String("stan"),
				},
			},
		}

		expectedBeer := &entities.Beer{
			ID: "stan",
		}

		mapped, _ := dynamodbattribute.MarshalMap(expectedBeer)
		m.On("GetItem", input).Return(&dynamodb.GetItemOutput{
			Item: mapped,
		}, nil)

		actual, err := repo.GetBeer("stan")
		assert.Nil(t, err)
		assert.Equal(t, expectedBeer, actual)
	})

	t.Run("it returns an error when the client errors", func(t *testing.T) {
		repo := setupFixture()
		m := repo.client.(*mocks.MockDynamoDBClient)

		testID := "stan"

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

		actual, err := repo.GetBeer("stan")

		assert.Nil(t, actual)
		assert.Equal(t, expErr, err)

		m.AssertExpectations(t)
	})
}

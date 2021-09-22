package breweries

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	testTable = "testTable"
	idField = "id"
)

func setupFixture() *Dynamo {
	return &Dynamo{
		client: &MockDynamoDBClient{},
	tableName: testTable,
	}
}

func Test_NewDynamo(t *testing.T) {
	t.Run("it returns a valid brewery dynamo", func(t *testing.T) {
		actual, err := NewDynamo(&DynamoConfig{
			Client: &MockDynamoDBClient{},
			TableName: testTable,
		})
		assert.Nil(t, err)
		assert.NotNil(t, actual)
		assert.NotNil(t, actual.client)
		assert.Equal(t, testTable, actual.tableName)
	})

	t.Run("it requires a brewery config", func(t *testing.T){
		actual, err := NewDynamo(nil)
		expErr := errors.New("it requires a brewery config")
		assert.Nil(t, actual)
		assert.NotNil(t, expErr, err)
	})

	t.Run("it requires a Client", func(t *testing.T){
		actual, err :=NewDynamo(&DynamoConfig{})
		expErr := "brewery cfg.Client is required"
		assert.Nil(t, actual)
		assert.NotNil(t, err)
		assert.Equal(t, expErr, err)
	})

	t.Run("it requires a TableName", func(t *testing.T){
		actual, err := NewDynamo(&DynamoConfig{
			Client: &MockDynamoDBClient{},
		})
		expErr := "brewery TableName is required"
		assert.Nil(t, actual)
		assert.NotNil(t, err)
		assert.Equal(t, expErr, err)
	})
}
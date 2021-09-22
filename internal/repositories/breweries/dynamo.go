package breweries

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/dammitbilly0ne/buzz-tracker/internal/entities"
)

type Dynamo struct {
	client    dynamodbiface.DynamoDBAPI
	tableName string
}

type DynamoConfig struct {
	Client    dynamodbiface.DynamoDBAPI
	TableName string
}

func NewDynamo(cfg *DynamoConfig) (*Dynamo, error) {
	if cfg == nil {
		return nil, errors.New("cfg is required")
	}
	if cfg.Client == nil {
		return nil, errors.New("Client is required")
	}
	if cfg.TableName == "" {
		return nil, errors.New("TableName is required")
	}

	return &Dynamo{
		client:    cfg.Client,
		tableName: cfg.TableName,
	}, nil
}

func (r *Dynamo) CreateBrewery(brewery *entities.Brewery) error {
	if brewery == nil {
		return errors.New("brewery is required")
	}
	if brewery.ID =="" {
		return errors.New("ID is required")
	}
	if brewery.Name == "" {
		return errors.New("brewery name is required")
	}

	return r.doCreateBrewery(brewery)
}

func (r *Dynamo) doCreateBrewery(brewery *entities.Brewery) error {
	mapped, err := dynamodbattribute.MarshalMap(brewery)
	if err != nil {
		return err
	}

	_, err = r.client.PutItem(&dynamodb.PutItemInput{
		TableName: &r.tableName,
		Item:      mapped,
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *Dynamo) FindBrewery(id string) (*entities.Brewery, error) {
	input := &dynamodb.GetItemInput{
		TableName: &r.tableName,
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	}
	output, err := r.client.GetItem(input)
	if err != nil {
		return nil, err
	}

	brewery := &entities.Brewery{}

	err = dynamodbattribute.UnmarshalMap(output.Item, brewery)
	if err != nil {
		return nil, err
	}

	return brewery, nil

}

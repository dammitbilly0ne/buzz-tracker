package beers

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/dammitbilly0ne/buzz-tracker/internal/entities"
)

const (
	beerRequiredMsg     = "beer is required"
	beerNameRequiredMsg = "name that beer"
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
		return nil, errors.New("cfg.Client is required")
	}
	if cfg.TableName == "" {
		return nil, errors.New("cfg.TableName is required")
	}

	return &Dynamo{
		client:    cfg.Client,
		tableName: cfg.TableName,
	}, nil
}

func (r *Dynamo) StoreBeer(beer *entities.Beer) error {
	if beer == nil {
		return errors.New(beerRequiredMsg)
	}
	if beer.Name == "" {
		return errors.New(beerNameRequiredMsg)
	}

	return r.doStoreBeer(beer)
}

func (r *Dynamo) GetBeer(id string) (*entities.Beer, error) {
	input := &dynamodb.GetItemInput{
		TableName: &r.tableName,
		Key: map[string]*dynamodb.AttributeValue{
			"id": &dynamodb.AttributeValue{
				S: aws.String(id),
			},
		},
	}

	output, err := r.client.GetItem(input)
	if err != nil {
		return nil, err
	}

	beer := &entities.Beer{}

	err = dynamodbattribute.UnmarshalMap(output.Item, beer)
	if err != nil {
		return nil, err
	}

	return beer, nil
}

func (r *Dynamo) doStoreBeer(beer *entities.Beer) error {
	mapped, err := dynamodbattribute.MarshalMap(beer)
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

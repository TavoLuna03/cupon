package repositories

import (
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/tavo/prueba/coupon/models"
)

const table = "items-test"

// ItemRepository intance to item repository
type ItemRepository struct {
	client *dynamodb.DynamoDB
}

// CreateItem create item to Dynamo
func (i *ItemRepository) CreateItem(id string, ammount float32) error {
	ammountItem := fmt.Sprintf("%f", ammount)
	_, err := i.client.PutItem(
		&dynamodb.PutItemInput{
			Item: map[string]*dynamodb.AttributeValue{
				"id": {
					S: aws.String(id),
				},
				"amount": {
					N: aws.String(ammountItem),
				},
			},
			TableName: aws.String(table),
		},
	)

	if err != nil {
		return err
	}
	return nil
}

// GetItems get all items
func (i *ItemRepository) GetItems() ([]models.Item, error) {
	params := &dynamodb.ScanInput{
		TableName: aws.String(table),
	}
	result, err := i.client.Scan(params)
	if err != nil {
		return []models.Item{}, err
	}

	elements, err := i.hydrate(result.Items)
	if err != nil {
		return []models.Item{}, err
	}
	return elements, nil
}

func (i *ItemRepository) hydrate(items []map[string]*dynamodb.AttributeValue) ([]models.Item, error) {
	itemsDynamo := make([]models.Item, len(items))
	for i, item := range items {
		itemsDynamo[i].ID = *item["id"].S

		if v, ok := item["amount"]; ok {
			value, err := strconv.ParseFloat(*v.N, 32)
			if err != nil {
				return []models.Item{}, err
			}
			itemsDynamo[i].Price = float32(value)
		}
	}
	return itemsDynamo, nil
}

// NewItemRepository trigger new repository
func NewItemRepository(client *dynamodb.DynamoDB) *ItemRepository {
	return &ItemRepository{
		client: client,
	}
}

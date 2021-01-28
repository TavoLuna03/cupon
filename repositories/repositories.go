package repositories

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/tavo/prueba/coupon/models"
)

// Repository instance repository
type Repository struct {
	client *http.Client
	host   string
}

// BodyItem entity
type BodyItem struct {
	Code int         `json:"code"`
	Body models.Item `json:"body"`
}

const itemURL = "/items"

//GetItemByID get price by id
func (i *Repository) GetItemByID(ids string) ([]models.Item, error) {

	response, err := i.client.Get(i.host + itemURL + "?ids=" + ids + "&attributes=id,price")
	if err != nil {
		return []models.Item{}, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		bodyErr, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return []models.Item{}, errors.New(string(bodyErr))
		}
	}

	var bodyItems []BodyItem
	byteValue, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []models.Item{}, err
	}

	err = json.Unmarshal(byteValue, &bodyItems)
	if err != nil {
		return []models.Item{}, err
	}

	var items []models.Item
	for _, itemRes := range bodyItems {
		item := models.Item{
			ID:    itemRes.Body.ID,
			Price: itemRes.Body.Price,
		}
		items = append(items, item)
	}
	return items, nil
}

// NewRepository initializer
func NewRepository(
	client *http.Client,
	host string,
) *Repository {
	return &Repository{
		client: client,
		host:   host,
	}
}

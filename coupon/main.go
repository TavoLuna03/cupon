package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/tavo/prueba/coupon/models"
	"github.com/tavo/prueba/coupon/usecases"
	"github.com/tavo/prueba/repositories"
)

type Handler func(ctx context.Context, req events.APIGatewayProxyRequest) (Response, error)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// RequestBody is a struct to request in the APIGatewayProxyRequest
type RequestBody struct {
	ItemIds []string `json:"item_ids"`
	Amount  int      `json:"amount"`
}

// Transport contract transport
type Transport struct {
	accept        string
	Authorization string
	rt            http.RoundTripper
}

func (t *Transport) transport() http.RoundTripper {
	if t.rt != nil {
		return t.rt
	}
	return http.DefaultTransport
}

// ResponseAPI contract response
type ResponseAPI struct {
	Items []string `json:"item_ids"`
	Total float64  `json:"total"`
}

// RoundTrip add headers
func (t *Transport) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Set("Accept", t.accept)
	return t.transport().RoundTrip(r)
}

// RepositoryInterface is a contract for a repository in this lambda
type RepositoryInterface interface {
	GetItemByID(ids string) ([]models.Item, error)
}

// Adapter  function call
func Adapter(repository RepositoryInterface) Handler {
	return func(ctx context.Context, req events.APIGatewayProxyRequest) (Response, error) {
		var requestBody RequestBody
		err := json.Unmarshal([]byte(req.Body), &requestBody)
		if err != nil {
			return Response{StatusCode: http.StatusNotFound}, err
		}

		ids := strings.Join(requestBody.ItemIds, ",")

		// get items Mercadolibre
		items, err := repository.GetItemByID(ids)
		if err != nil {
			return Response{StatusCode: http.StatusNotFound}, err
		}

		var buf bytes.Buffer

		finalSolution := make([]string, 0)
		usecase := usecases.NewUseCases([]models.Item{}, []models.Item{}, finalSolution)
		fmt.Println(requestBody.Amount)
		err = usecase.ValidatePriceMin(float64(requestBody.Amount), items)
		if err != nil {
			return Response{
				StatusCode:      http.StatusNotFound,
				IsBase64Encoded: false,
				Body:            err.Error(),
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
			}, nil
		}
		response := usecase.Calculate(items, float64(requestBody.Amount))
		finalItems := usecase.GetItemWithPrice(response, items)
		total := usecase.CalculateTotal(finalItems)

		responseService := ResponseAPI{
			Items: response,
			Total: total,
		}

		body, err := json.Marshal(responseService)
		if err != nil {
			return Response{StatusCode: http.StatusNotFound}, err
		}
		json.HTMLEscape(&buf, body)
		resp := Response{
			StatusCode:      http.StatusOK,
			IsBase64Encoded: false,
			Body:            buf.String(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return resp, nil
	}
}

func main() {
	token := os.Getenv("TOKEN")
	host := os.Getenv("HOST")

	fmt.Printf("asdfsdaf %v", token)
	client := http.Client{
		Timeout: time.Second * time.Duration(30),
	}
	client.Transport = &Transport{
		accept:        "application/json",
		Authorization: "Bearer " + token,
		rt:            client.Transport,
	}
	repository := repositories.NewRepository(&client, host)
	lambda.Start(Adapter(repository))
}

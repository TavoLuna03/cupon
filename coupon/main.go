package main

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

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

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Adapter() Handler {
	return func(ctx context.Context, req events.APIGatewayProxyRequest) (Response, error) {
		var requestBody RequestBody
		err := json.Unmarshal([]byte(req.Body), &requestBody)
		if err != nil {
			return Response{StatusCode: 404}, err
		}

		sess := session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		}))

		// Create DynamoDB client
		svc := dynamodb.New(sess)
		repository := repositories.NewItemRepository(svc)
		itemsDynamo, err := repository.GetItems()
		if err != nil {
			return Response{
				StatusCode:      200,
				IsBase64Encoded: false,
				Body:            "error",
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
			}, nil
		}

		var buf bytes.Buffer

		finalSolution := make([]string, 0)
		usecase := usecases.NewUseCases([]models.Item{}, []models.Item{}, finalSolution)

		items := usecase.GetItemWithPrice(requestBody.ItemIds, itemsDynamo)
		err = usecase.ValidatePriceMin(float32(requestBody.Amount), items)
		if err != nil {
			return Response{
				StatusCode:      404,
				IsBase64Encoded: false,
				Body:            err.Error(),
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
			}, nil
		}

		response := usecase.Calculate(items, float32(requestBody.Amount))
		body, err := json.Marshal(response)
		if err != nil {
			return Response{
				StatusCode:      404,
				IsBase64Encoded: false,
				Body:            err.Error(),
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
			}, nil
		}

		json.HTMLEscape(&buf, body)

		resp := Response{
			StatusCode:      200,
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
	lambda.Start(Adapter())
}

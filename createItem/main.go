package main

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

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
	ID     string `json:"id"`
	Amount int    `json:"amount"`
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
		err = repository.CreateItem(requestBody.ID, float32(requestBody.Amount))
		if err != nil {
			return Response{StatusCode: 404}, err
		}
		var buf bytes.Buffer

		itemsDynamo, err := repository.GetItems()
		if err != nil {
			return Response{StatusCode: 404}, err
		}
		body, err := json.Marshal(itemsDynamo)
		if err != nil {
			return Response{StatusCode: 404}, err
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

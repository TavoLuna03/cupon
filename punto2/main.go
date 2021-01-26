package main

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/tavo/prueba/punto2/models"
	"github.com/tavo/prueba/punto2/usecases"
)

type Handler func(ctx context.Context) (Response, error)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Adapter() Handler {
	return func(ctx context.Context) (Response, error) {
		var buf bytes.Buffer
		a := []models.Item{
			{
				ID:    "MLA1",
				Price: 100,
			},
			{
				ID:    "MLA2",
				Price: 210,
			},
			{
				ID:    "MLA3",
				Price: 260,
			},
			{
				ID:    "MLA4",
				Price: 80,
			},
			{
				ID:    "MLA5",
				Price: 90,
			},
		}

		optima := []models.Item{}
		solutionItems := []models.Item{}
		finalSolution := make([]string, 0)
		usecase := usecases.NewUseCases(optima, solutionItems, finalSolution)
		response := usecase.Calculate(a, 500)
		body, err := json.Marshal(map[string]interface{}{
			"message": response,
		})
		if err != nil {
			return Response{StatusCode: 404}, err
		}

		json.HTMLEscape(&buf, body)

		resp := Response{
			StatusCode:      200,
			IsBase64Encoded: false,
			Body:            buf.String(),
			Headers: map[string]string{
				"Content-Type":           "application/json",
				"X-MyCompany-Func-Reply": "hello-handler",
			},
		}
		return resp, nil
	}
}

func main() {

	lambda.Start(Adapter())
}

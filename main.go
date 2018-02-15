package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Body size = %d. \n", len(request.Body))
	fmt.Println("Headers:")
	for key, value := range request.Headers {
		fmt.Printf("  %s: %s\n", key, value)
	}
	if request.HTTPMethod == "GET" {
		fmt.Printf("GET METHOD\n")
		return events.APIGatewayProxyResponse{Body: "GET", StatusCode: 200}, nil
	} else if request.HTTPMethod == "POST" {
		fmt.Printf("POST METHOD\n")
		return events.APIGatewayProxyResponse{Body: "POST", StatusCode: 200}, nil
	} else {
		fmt.Printf("NEITHER\n")
		return events.APIGatewayProxyResponse{Body: "NEITHER", StatusCode: 200}, nil
	}
}

func main() {
	lambda.Start(HandleRequest)
}

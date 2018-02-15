package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	// ErrNameNotProvided is thrown when a name is not provided
	ErrNameNotProvided = errors.New("no name was provided in the HTTP body")
)

// var ddb *dynamodb.DynamDB

// func init() {
// 	region := os.Getenv("AWS_REGION")
// 	if session, err := session.NewSession(&aws.Config{ // Use aws sdk to connect to dynamoDB
// 		Region: &region,
// 	}); err != nil {
// 		fmt.Println(fmt.Sprintf("Failed to connect to AWS: %s", err.Error()))
// 	} else {
// 		ddb = dynamodb.New(session) // Create DynamoDB client
// 	}
// }

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Body size = %d. \n", len(request.Body))
	fmt.Println("Headers:")
	for key, value := range request.Headers {
		fmt.Printf("  %s: %s\n", key, value)
	}
	if request.HTTPMethod == "GET" {
		fmt.Printf("GET METHOD")
	} else if request.HTTPMethod == "POST" {
		fmt.Printf("POST METHOD")
	} else {
		fmt.Printf("NEITHER!")
	}
	return events.APIGatewayProxyResponse{Body: "Success!", StatusCode: 200}, nil
}

func main() {
	lambda.Start(handleRequest)
}

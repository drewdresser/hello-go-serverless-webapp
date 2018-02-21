package main_test

import (
	"testing"

	main "github.com/drewdresser/hello-go-serverless-webapp"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {

	tests := []struct {
		request events.APIGatewayProxyRequest
		expect  string
		err     error
	}{
		{
			// Test that the handler responds with the correct response
			// when a valid name is provided in the HTTP body
			request: events.APIGatewayProxyRequest{HTTPMethod: "GET"},
			expect:  "GET",
			err:     nil,
		},
		{
			// Test that the handler responds ErrNameNotProvided
			// when no name is provided in the HTTP body
			request: events.APIGatewayProxyRequest{Body: "PUT"},
			expect:  "",
			err:     main.HTTPMethodNotSupported,
		},
	}

	for _, test := range tests {
		response, err := main.HandleRequest(nil, test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.expect, response.Body)
	}

}

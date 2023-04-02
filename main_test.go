package main

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {

	request := events.APIGatewayProxyRequest{
		QueryStringParameters: map[string]string{
			"name": "world",
		},
	}
	expectedResponse := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "text/html",
		},
		Body: "Hello, world",
	}

	response, err := Handler(request)

	assert.Equal(t, response.Headers, expectedResponse.Headers)
	assert.Contains(t, response.Body, expectedResponse.Body)
	assert.Equal(t, err, nil)

}

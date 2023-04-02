package main

import (
	"fmt"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {

	request := events.APIGatewayProxyRequest{}

	response, err := Handler(request)

	fmt.Println(response)

	assert.Equal(t, response.Headers, map[string]string{
		"Content-Type": "text/html",
	})
	assert.Equal(t, err, nil)

	assert.NotContains(t, response.Body, "You can buy")

	assert.Contains(t, response.Body, "Enter the amount of dollars to convert to bitcoin")

}

package main

import (
    "fmt"
    "net/http"

    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    name := request.QueryStringParameters["name"]
    if name == "" {
        return events.APIGatewayProxyResponse{
            StatusCode: http.StatusOK,
            Body:       "<html><body><form><label>Enter your name:</label><input type=\"text\" name=\"name\"><input type=\"submit\" value=\"Submit\"></form></body></html>",
        }, nil
    } else {
        message := fmt.Sprintf("<html><body><h1>Hello, %s!</h1><form><label>Enter your name:</label><input type=\"text\" name=\"name\" value=\"%s\"><input type=\"submit\" value=\"Submit\"></form></body></html>", name, name)
        return events.APIGatewayProxyResponse{
            StatusCode: http.StatusOK,
            Body:       message,
        }, nil
    }
}

func main() {
    lambda.Start(handler)
}

package main

import (
	"bytes"
	"html/template"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)


func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    var buf bytes.Buffer

    index, err1 := template.ParseFiles("public/index.html")
    if err1 != nil {
        return events.APIGatewayProxyResponse{
            Body: err1.Error(),
        }, err1
    }

    err2 := index.Execute(&buf, request.QueryStringParameters)
    
    if err2 != nil {
        return events.APIGatewayProxyResponse{
            Body: err2.Error(),
        }, err2
    }

    return events.APIGatewayProxyResponse{
        StatusCode: 200,
        Body:       buf.String(),
        Headers: map[string]string{
            "Content-Type": "text/html",
        },
    }, nil

}

func main() {
	lambda.Start(Handler)
}

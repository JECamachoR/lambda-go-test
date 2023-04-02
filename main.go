package main

import (
	"bytes"
	"encoding/json"
	"html/template"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"gopkg.in/resty.v1"
)

// ExchangeRateResponse represents the response from the CoinAPI exchange rate API
type ExchangeRateResponse struct {
	Time         string  `json:"time"`
	AssetIDBase  string  `json:"asset_id_base"`
	AssetIDQuote string  `json:"asset_id_quote"`
	Rate         float64 `json:"rate"`
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
// it takes one query parameter of `usd` and returns the equivalent in BTC at the current exchange rate
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var usd float64
	if usdStr, ok := request.QueryStringParameters["USD"]; ok {
		var err error
		usd, err = strconv.ParseFloat(usdStr, 64)
		if err != nil {
			return events.APIGatewayProxyResponse{
				Body: err.Error(),
			}, err
		}
	}

	var btc float64
	if usd != 0 {
		client := resty.New()

		resp, err := client.R().
			SetHeader("X-CoinAPI-Key", "6E3C5583-2C28-45DF-85C6-2338FDAA965B").
			Get("https://rest.coinapi.io/v1/exchangerate/BTC/USD")

		if err != nil {
			return events.APIGatewayProxyResponse{
				Body: err.Error(),
			}, err
		}

		var exchangeRate ExchangeRateResponse
		err = json.Unmarshal(resp.Body(), &exchangeRate)
		if err != nil {
			return events.APIGatewayProxyResponse{
				Body: err.Error(),
			}, err
		}

		btcPrice := exchangeRate.Rate

		btc = usd / btcPrice
	}

	var buf bytes.Buffer

	index, err := template.ParseFiles("public/index.html")
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body: err.Error(),
		}, err
	}

	index.Execute(&buf, struct {
		USD float64
		BTC float64
	}{
		USD: usd,
		BTC: btc,
	})

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

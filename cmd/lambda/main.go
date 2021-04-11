package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mtanzim/guac/server/services"
	"github.com/mtanzim/guac/server/utils"
)

type MyEvent struct {
	QueryStringParameters struct {
		Start string `json:"start"`
		End   string `json:"end"`
	} `json:"queryStringParameters"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	start := request.QueryStringParameters["start"]
	end := request.QueryStringParameters["end"]

	if err := utils.ValidateQueryDate(start, end); err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	rv := services.DataService(start, end)
	b, err := json.MarshalIndent(rv, "", "  ")
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	return events.APIGatewayProxyResponse{
		Body:       string(b),
		StatusCode: 200,
	}, nil

}

func main() {
	lambda.Start(handler)
}

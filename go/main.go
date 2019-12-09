package main

import (
	"log"

	"github.com/kgrvamsi/aws-lambda-examples/go/weather"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	log.Println(request.RequestContext)
	return events.APIGatewayProxyResponse{
		Body:       weather.AdvancedHelloWorld(),
		StatusCode: 200,
	}, nil

}

func main() {

	lambda.Start(Handler)
}

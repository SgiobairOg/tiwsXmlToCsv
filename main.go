package main

import (
	"errors"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	// ErrNameNotProvided is thrown when name is not provided
	ErrNameNotProvided = errors.New("no name was provided in the HTTP body")
)

// Handler for Lambda function
// Use Amazon API Gateway request/response from the aws-lambda-go/events package
// Could use other event sources (S3, Kinesis, etc), or JSON-decoded primitives
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// Output and errors go to AWS Cloudwatch Logs
	log.Printf("Processing Lambda request %s\n", request.RequestContext.RequestID)

	if len(request.Body) < 1 {
		return events.APIGatewayProxyResponse{}, ErrNameNotProvided
	}

	return events.APIGatewayProxyResponse{
		Body:       "Hello " + request.Body,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}

package utils

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/pkg/errors"
)

func ErrorResponse(err error, message string) (events.APIGatewayProxyResponse, error) {
	err = errors.Wrap(err, message)
	log.Println(err)
	errStr := err.Error()
	return makeResponse(&errStr, 500)
}

func UnauthorizedResponse(body *string) (events.APIGatewayProxyResponse, error) {
	return makeResponse(body, 401)
}

func OkResponse(body *string) (events.APIGatewayProxyResponse, error) {
	return makeResponse(body, 200)
}

func CreatedResponse(body *string) (events.APIGatewayProxyResponse, error) {
	return makeResponse(body, 201)
}

func BadRequestResponse(body *string) (events.APIGatewayProxyResponse, error) {
	return makeResponse(body, 400)
}

func makeResponse(body *string, statusCode int) (events.APIGatewayProxyResponse, error) {
	res := events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-Type":                 "application/json",
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Methods": "*",
			"Access-Control-Allow-Headers": "*",
		},
		StatusCode: statusCode,
	}

	if body != nil {
		res.Body = *body
	}

	return res, nil
}

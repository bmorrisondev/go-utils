package utils

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/pkg/errors"
)

type LambdaRouter struct {
	Get    func(events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
	Post   func(events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
	Put    func(events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
	Patch  func(events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
	Delete func(events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
}

func (lr *LambdaRouter) Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if request.HTTPMethod == "GET" && lr.Get != nil {
		return lr.Get(request)
	}
	if request.HTTPMethod == "POST" && lr.Post != nil {
		return lr.Post(request)
	}
	if request.HTTPMethod == "PUT" && lr.Put != nil {
		return lr.Put(request)
	}
	if request.HTTPMethod == "PATCH" && lr.Patch != nil {
		return lr.Patch(request)
	}
	if request.HTTPMethod == "DELETE" && lr.Delete != nil {
		return lr.Delete(request)
	}
	return NotFoundResponse()
}

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

func NotFoundResponse() (events.APIGatewayProxyResponse, error) {
	return makeResponse(nil, 400)
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

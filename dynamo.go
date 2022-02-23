package utils

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/pkg/errors"
)

type DynamoContext struct {
	DynamoSvc *dynamodb.DynamoDB
	Session   *session.Session
	TableName *string
}

func NewDynamoContext(tableName string, sess *session.Session) (*DynamoContext, error) {
	if tableName == "" {
		return nil, errors.New("(NewDynamoContext) tableName is required")
	}
	var _session *session.Session
	if sess == nil {
		sess, err := session.NewSession()
		if err != nil {
			return nil, errors.Wrap(err, "(MakeContext) creating session")
		}
		_session = sess
	} else {
		_session = sess
	}

	svc := dynamodb.New(_session)

	context := DynamoContext{
		DynamoSvc: svc,
		Session:   _session,
	}

	context.TableName = &tableName
	return &context, nil
}

func (context *DynamoContext) Put(obj interface{}) error {
	item, err := dynamodbattribute.MarshalMap(obj)
	if err != nil {
		return errors.Wrap(err, "(PutBuildToDynamo) Marshal map")
	}

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: context.TableName,
	}

	_, err = context.DynamoSvc.PutItem(input)
	if err != nil {
		return err
	}

	return nil
}

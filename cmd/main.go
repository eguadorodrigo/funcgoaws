package cmd

import (
	"fmt"
	"os"
	"net/http"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/eguadorodrigo/funcgoaws/handlers"
)
func main()(
	region := os.Getenv("AWS_REGION")
	awsSession, err := session.NewSession(&aws.Config{
		Region := aws.String(region)},)
		if err != nil{
			return
		}
	dynaClient = dynamodb.New(awsSession)
	lambda.Start(handler)	
)

const tableName = "lambda-table-go"

func handler(req events.APIGatewayProxyRequest)(*events.APIGatewayProxyResponse, error){
	switch req.HTTPMethod{
		case http.MethodGet:
			return handlers.GetUser(req, tableName, dynaClient)
		case http.MethodPost:
			handlers.CreateUser(req, tableName, dynaClient)
		case http.MethodDelete:
			handlers.DeleteUser(req, tableName, dynaClient)
		case http.MethodPut:
			handlers.UpdateUser(req, tableName, dynaClient)
		default:
			return handlers.UnhandlerMethod()
	}

}
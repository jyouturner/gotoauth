package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/jyouturner/gotoauth"
	"github.com/jyouturner/gotoauth/example/awsserverless"
	lambdahelper "github.com/jyouturner/gotoauth/example/awsserverless/lambda"
)

func init() {

	logLevel, exists := os.LookupEnv("LOG_LEVEL")
	if exists {
		level, err := log.ParseLevel(logLevel)
		if err != nil {
			log.Errorf("incorrect LOG_LEVEL %s", level)
		} else {
			log.SetLevel(level)
		}
	}
	_, exists = os.LookupEnv("AWS_SECRET_NAME")
	if !exists {
		log.Fatalf("missing %s", "AWS_SECRET_NAME")
	}
	_, exists = os.LookupEnv("ACCESS_TOKEN_BUCKET")
	if !exists {
		log.Fatalf("missing %s", "ACCESS_TOKEN_BUCKET")
	}
	_, exists = os.LookupEnv("OAUTH_NOUNCE_BUCKET")
	if !exists {
		log.Fatalf("missing %s", "OAUTH_NOUNCE_BUCKET")
	}
	_, exists = os.LookupEnv("OAUTH_PROVIDER")
	if !exists {
		log.Fatalf("missing %s", "OAUTH_PROVIDER")
	}
	_, exists = os.LookupEnv("SCOPE")
	if !exists {
		log.Fatalf("missing %s", "SCOPE")
	}
	_, exists = os.LookupEnv("AUTHORIZED_TO_URL")
	if !exists {
		log.Fatalf("missing %s", "AUTHORIZED_TO_URL")
	}

}

func main() {
	lambda.Start(Handle)
}

func Handle(ctx context.Context, event json.RawMessage) (lambdahelper.LambdaResponse, error) {
	fmt.Println(string(event))
	eventBodyString := gjson.Get(string(event), "body").String()
	fmt.Println(string(eventBodyString))
	//get user data from body
	/*
		{
		"user": {
			"org_id": "12345",
			"user_id": "abcde"
		}
	*/
	userData := gjson.Get(eventBodyString, "user").String()
	orgUser, err := awsserverless.FromJson([]byte(userData))
	if err != nil {
		log.Errorf("failed to convert json to user meta %v %v", userData, err)
		return lambdahelper.FailureMessage(400, "user data is not expected"), err
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return lambdahelper.FailureMessage(500, "system error"), fmt.Errorf("system error")

	}

	awsClient := awsserverless.AWSClient{
		Config: cfg,
	}

	nounceState := awsserverless.StateToken{
		User:               orgUser,
		Provider:           strings.ToUpper(os.Getenv("OAUTH_PROVIDER")),
		Scope:              os.Getenv("SCOPE"),
		SuccessRedirectUrl: os.Getenv("AUTHORIZED_TO_URL"),
	}
	awsEnv, err := awsserverless.NewAWSEnvByUser(awsClient, os.Getenv("AWS_SECRET_NAME"), os.Getenv("ACCESS_TOKEN_BUCKET"), orgUser, os.Getenv("OAUTH_NOUNCE_BUCKET"))
	if err != nil {
		return lambdahelper.FailureMessage(500, "could not load oauth config from aws"), err
	}

	authurl, err := gotoauth.GetAuthUrl(nounceState, awsEnv, awsEnv)

	if err != nil {
		log.Errorf("failed to initialize authorization flow for user %v %v", orgUser, err)
		return lambdahelper.FailureMessage(500, "failed to get auth url"), err
	}
	return lambdahelper.Success(*authurl), nil
}

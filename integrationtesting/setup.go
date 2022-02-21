package integrationtesting

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/jyouturner/gotoauth/awssolution"
)

func getAWSClient(profile string, t *testing.T) awssolution.AWSClient {

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithSharedConfigProfile(profile))
	if err != nil {
		t.Errorf("failed to create aws session, %v", err)

	}
	awsClient := awssolution.AWSClient{
		Config: cfg,
	}
	if err != nil {
		t.Error("could not create aws client")
	}
	return awsClient
}

func getTestData(file string, t *testing.T) map[string]string {
	//load the testing json data
	b, err := ioutil.ReadFile(file)
	if err != nil {
		t.Errorf("failed to load the tesing json from %v", err)
	}
	data := map[string]string{}
	err = json.Unmarshal(b, &data)
	if err != nil {
		t.Errorf("failed to parse testing json data %v", err)
	}
	return data
}

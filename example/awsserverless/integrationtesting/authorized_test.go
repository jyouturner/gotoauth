package integrationtesting

import (
	"testing"

	"github.com/jyouturner/gotoauth"
	"github.com/jyouturner/gotoauth/example/awsserverless"
)

func TestExchangeTokenWithGoogle(t *testing.T) {
	data := getTestData("testdata/TestExchangeTokenWithGoogle.json", t)
	awsClient := getAWSClient(data["TEST_AWS_PROFILE"], t)

	awsEnv, err := awsserverless.NewAWSEnv(awsClient, data["AWS_SECRET_NAME"], data["TOKEN_BUCKET"], data["USER_ID"], data["NOUNCE_TOKEN_BUCKET"])
	if err != nil {
		t.Errorf("could not load oauth config from aws %v", err)
	}

	authcode := ""
	state := ""
	if err := gotoauth.Exchange(authcode, state, awsEnv, awsEnv); err != nil {
		t.Errorf("Exchange() error = %v", err)
	}

}

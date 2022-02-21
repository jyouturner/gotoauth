package integrationtesting

import (
	"os"
	"testing"

	"github.com/jyouturner/gotoauth"
	"github.com/jyouturner/gotoauth/awssolution"
	log "github.com/sirupsen/logrus"
)

func TestGetAuthUrlGoogle(t *testing.T) {
	logLevel, exists := os.LookupEnv("LOG_LEVEL")
	if exists {
		level, err := log.ParseLevel(logLevel)
		if err != nil {
			log.Errorf("incorrect LOG_LEVEL %s", level)
		} else {
			log.SetLevel(level)
		}
	}
	//load the testing json data
	data := getTestData("testdata/TestGetAuthUrlGoogle.json", t)
	awsClient := getAWSClient(data["TEST_AWS_PROFILE"], t)
	provider := data["PROVIDER"]
	user := awssolution.UserMeta{
		OrgId:  data["ORG_ID"],
		UserId: data["USER_ID"],
	}
	nounceState := awssolution.StateToken{
		User:     user,
		Provider: provider,
		Scope:    data["SCOPE"],
	}
	awsEnv, err := awssolution.NewAWSEnvByUser(awsClient, data["AWS_SECRET_NAME"], data["TOKEN_BUCKET"], user, data["NOUNCE_TOKEN_BUCKET"])
	if err != nil {
		t.Errorf("could not load oauth config from aws %v", err)
	}
	url, err := gotoauth.GetAuthUrl(nounceState, awsEnv, awsEnv)
	if err != nil {
		t.Errorf("could not get auth url from %s %v", provider, err)
	}
	t.Log(*url)

}

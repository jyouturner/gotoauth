//when the app is deployed at AWS, we will use below to save the environment and configurations. This is the SaaS style.
//aws secret manager to store the google and atlanssian secret JSON. In this case, we need the secret name, which can have multiple key-value pairs.
//the oauth token will be saved in S3, under the user's folder, the bucket will be the organization bucket.
//the user config (config.yml) will be saved in the S3 as well, under user folder.
//the S3 bucket structure
// tam-auth-state
// tam-org-12345
//		/abcde
//			/google_token.json
//			/config.yml
package integrationtesting

import (
	"testing"

	"github.com/jyouturner/gotoauth/example/awsserverless"
)

func TestNewAwsEnv(t *testing.T) {
	data := getTestData("testdata/AWSEnvTest.json", t)
	awsClient := getAWSClient(data["TEST_AWS_PROFILE"], t)
	awsEnv, err := awsserverless.NewAWSEnv(awsClient, data["AWS_SECRET_NAME"], data["TOKEN_BUCKET"], data["USER_ID"], data["NOUNCE_TOKEN_BUCKET"])
	if err != nil {
		t.Errorf("could not load oauth config from aws %v", err)
	}
	_, err = awsEnv.GetAppOathConfig("GOOGLE")
	if err != nil {
		t.Errorf("missing auth config for provider")
	}

}

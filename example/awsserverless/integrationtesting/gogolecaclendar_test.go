package integrationtesting

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/jyouturner/gotoauth"
	"github.com/jyouturner/gotoauth/example/awsserverless"
	log "github.com/sirupsen/logrus"
	googlecalendar "google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

func TestListGoogleCalendarEvents(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	data := getTestData("testdata/TestListGoogleCalendarEvents.json", t)
	awsClient := getAWSClient(data["TEST_AWS_PROFILE"], t)
	user := awsserverless.OrgUser{
		OrgId:  data["ORG_ID"],
		UserId: data["USER_ID"],
	}
	awsEnv, err := awsserverless.NewAWSEnvByUser(awsClient, data["AWS_SECRET_NAME"], data["TOKEN_BUCKET"], user, data["NOUNCE_TOKEN_BUCKET"])
	if err != nil {
		t.Errorf("could not load oauth config from aws %v", err)
	}
	authconfig, err := awsEnv.GetAppOathConfig("GOOGLE")
	if err != nil {
		t.Errorf("missing auth config for provider")
	}
	oauthConfig, err := gotoauth.ConfigFromJSON(authconfig.Secret, strings.Split(data["SCOPE"], " "))
	if err != nil {
		log.Fatalf("error get oauth config %v", err)
	}
	httpClient, err := gotoauth.NewClient(authconfig.OauthTokenStorage, oauthConfig)

	if err != nil {
		log.Fatalf("error create http client %v", err)
	}
	//query google calendar
	srv, err := googlecalendar.NewService(context.Background(), option.WithHTTPClient(httpClient))
	if err != nil {
		t.Errorf("unable to create Google Calendar Service: %v", err)
	}
	nowTime := time.Now().Format(time.RFC3339)
	_, err = srv.Events.List("Primary").ShowDeleted(false).
		SingleEvents(true).TimeMin(nowTime).MaxResults(10).OrderBy("startTime").Do()
	if err != nil {
		t.Errorf("could not fetch the calendar event %v", err)
	}
}

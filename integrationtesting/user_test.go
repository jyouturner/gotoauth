package integrationtesting

import (
	"fmt"
	"testing"

	"github.com/jyouturner/gotoauth/awssolution"
)

func TestUserMetaFromJson(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test json unmarshal",
			args: args{
				data: []byte("{\"org_id\": \"12345\",\"user_id\": \"abcde\"}"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := awssolution.UserMetaFromJson(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserMetaFromJson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)
		})
	}
}

func TestUserMetaFromJson2(t *testing.T) {
	eventBody := `
		{
		  "org_id": "12345",
		  "user_id": "abcde"
		}
	`
	userMeta, err := awssolution.UserMetaFromJson([]byte(eventBody))
	if err != nil {
		t.Error(err)
	}
	fmt.Println(userMeta.OrgId)
}

/*
func TestParseEventJson(t *testing.T) {
	eventBodyString := `
	{
		"user": {
		  "org_id": "12345",
		  "user_id": "abcde"
		},
		"oauth_provider": "Google"
	  }
	`

	authProvider := gjson.Get(eventBodyString, "oauth_provider").String()
	fmt.Println(authProvider)
	eventBody := gjson.Parse(eventBodyString)
	user := eventBody.Get("user").String()
	fmt.Println(user)
	um, err := UserMetaFromJson([]byte(user))

	if err != nil {
		t.Error(err)
	}

	if um.OrgId != "12345" {
		t.Errorf("expected 12345 get %v", um.OrgId)
	}
}
*/
func TestDecodeUserMeta(t *testing.T) {
	type args struct {
		data []byte
		um   *awssolution.UserMeta
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test decode user",
			args: args{
				data: []byte(`{"org_id":"12345","user_id":"abcde"}`),
				um:   &awssolution.UserMeta{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := awssolution.DecodeUserMeta(tt.args.data, tt.args.um); (err != nil) != tt.wantErr {
				t.Errorf("DecodeUserMeta() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
		fmt.Println(tt.args.um.OrgId)
	}
}

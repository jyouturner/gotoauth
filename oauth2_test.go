package gotoauth

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestOauth2_GetAccessibleResources(t *testing.T) {
	type fields struct {
		TokenFile  string
		SecretFile string
		Scope      []string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test oauth2",
			fields: fields{
				TokenFile:  "",
				SecretFile: "",
				Scope:      []string{"offline_access", "read:jira-user", "read:jira-work"},
			},
			want:    "ok",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//create the http client
			localTokenStorage := LocalTokenStorage{
				TokenFile: tt.fields.TokenFile,
			}
			config, _ := ConfigFromLocalJsonFile(tt.fields.SecretFile, tt.fields.Scope)
			client, err := NewClient(localTokenStorage, config)
			if err != nil {
				t.Fail()
			}
			log.Info(client)
		})
	}
}

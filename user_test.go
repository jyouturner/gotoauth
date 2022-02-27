package gotoauth

import (
	"fmt"
	"testing"
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
			got, err := FromJson(tt.args.data)
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
	userMeta, err := FromJson([]byte(eventBody))
	if err != nil {
		t.Error(err)
	}
	fmt.Println(userMeta)
}

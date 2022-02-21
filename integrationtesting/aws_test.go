package integrationtesting

import (
	"encoding/json"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestS3_save(t *testing.T) {
	testdata := getTestData("testdata/AWSTest.json", t)
	type fields struct {
	}
	type args struct {
		Bucket string
		Key    string
		data   []byte
	}
	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "test save file to s3",
			fields: fields{},
			args: args{

				data:   mapB,
				Bucket: testdata["BUCKET"],
				Key:    testdata["FILE"],
			},
			wantErr: false,
		},
	}
	client := getAWSClient(testdata["TEST_AWS_PROFILE"], t)
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			if err := client.S3Save(tt.args.data, tt.args.Bucket, tt.args.Key); (err != nil) != tt.wantErr {
				t.Errorf("S3TokenStorage.save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestS3TokenStorage_ls(t *testing.T) {
	testdata := getTestData("testdata/AWSTest.json", t)
	type fields struct {
	}
	type args struct {
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "test access",
			fields:  fields{},
			args:    args{},
			wantErr: false,
		},
	}
	client := getAWSClient(testdata["TEST_AWS_PROFILE"], t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := client.S3Ls()
			if err != nil && true != tt.wantErr {
				t.Errorf("S3TokenStorage.ls() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_get(t *testing.T) {
	testdata := getTestData("testdata/AWSTest.json", t)
	type args struct {
		bucket string
		key    string
	}

	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test get object",
			args: args{

				bucket: testdata["BUCKET"],
				key:    testdata["FILE"],
			},
			want:    5,
			wantErr: false,
		},
	}
	client := getAWSClient(testdata["TEST_AWS_PROFILE"], t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.S3Get(tt.args.bucket, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//mapD := map[string]int{"apple": 5, "lettuce": 7}
			data := map[string]int{}
			json.Unmarshal(got, &data)
			log.Info(data["apple"])
			if data["apple"] != tt.want {
				t.Errorf("get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSecret(t *testing.T) {
	testdata := getTestData("testdata/AWSTest.json", t)
	type args struct {
		secretName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test aws secret manager",
			args: args{
				secretName: testdata["SECRET"],
			},
			wantErr: false,
		},
	}
	client := getAWSClient(testdata["TEST_AWS_PROFILE"], t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := client.GetSecret(tt.args.secretName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSecret() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func TestSecretToEnvVariables(t *testing.T) {
	testdata := getTestData("testdata/AWSTest.json", t)
	type args struct {
		secretName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test load secret into environments",
			args: args{
				secretName: testdata["SECRET"],
			},
			wantErr: false,
		},
	}
	client := getAWSClient(testdata["TEST_AWS_PROFILE"], t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := client.SecretToEnvVariables(tt.args.secretName); (err != nil) != tt.wantErr {
				t.Errorf("SecretToEnvVariables() error = %v, wantErr %v", err, tt.wantErr)
			}

		})
	}
}

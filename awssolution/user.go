package awssolution

import (
	"encoding/json"
)

type UserMeta struct {
	OrgId  string `json:"org_id"`
	UserId string `json:"user_id"`
}

func DecodeUserMeta(data []byte, um *UserMeta) error {
	return json.Unmarshal(data, um)
}

func EncodeUserMeta(um UserMeta) ([]byte, error) {
	return json.Marshal(um)
}

func UserMetaFromJson(data []byte) (UserMeta, error) {
	um := UserMeta{}
	err := json.Unmarshal(data, &um)
	return um, err
}

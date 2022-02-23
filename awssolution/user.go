package awssolution

import (
	"encoding/json"
)

//OrgUser implements the UserMeta with organization and user identifier. This is probably enough for most of use cases.
type OrgUser struct {
	OrgId  string `json:"org_id"`
	UserId string `json:"user_id"`
}

func FromJson(data []byte) (UserMeta, error) {
	um := OrgUser{}
	err := json.Unmarshal(data, &um)
	return um, err
}

func (p OrgUser) Encode() ([]byte, error) {
	return json.Marshal(p)
}

func (p OrgUser) GetAccessTokenFolderPath() string {
	return p.UserId
}

package gotoauth

import (
	"fmt"
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

func init() {

}

//GetAuthUrl returns the oauth2 url to get autocode from the provider. It will create a random nounce and use it as key to store the "state" (which wraps the user's identifier, auth provider, whatnot).
func GetAuthUrl(state OauthState, configProvider OAuthConfigProvider, nounceWriter OauthNounceStateWriter) (*string, error) {
	log.Debugf("trying to find auth provider %s", state.GetProvider())
	authConfig, err := configProvider.GetOauthConfigOfUser(state.GetProvider())
	if err != nil {
		return nil, fmt.Errorf("failed to get the auth config of provider %s %v", state.GetProvider(), err)
	}

	cfg, err := ConfigFromJSON(authConfig.Secret, state.GetScope())

	if err != nil {
		return nil, fmt.Errorf("could not get the oauth config of the provider %v", err)
	}

	if cfg == nil {
		return nil, fmt.Errorf("no secret found matching the name")
	}

	nounce := String(16)

	//save the nounce for later verification
	err = nounceWriter.Save(nounce, state)
	if err != nil {
		return nil, fmt.Errorf("failed to save nounce %v", err)
	}

	//get the auth URL
	authUrl := cfg.AuthCodeURL(nounce, oauth2.AccessTypeOffline)

	//redirect to the oauth provider authrization URL
	return &authUrl, nil

}

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return StringWithCharset(length, charset)
}

package gotoauth

import (
	"context"
	"fmt"
)

func init() {

}

//Exchange gets the oauth2 access token from the auth code and save it, based on the oauth config.
func Exchange(authcode string, state string, configProvider OAuthConfigProvider, nounceReader OauthNounceStateReader) error {

	//find the user data from the nounce(state)
	stateTokenData, err := nounceReader.Read(state)
	if err != nil {
		return fmt.Errorf("failed to locate auth state data %v", err)
	}
	if stateTokenData == nil {
		return fmt.Errorf("no matching user found with the given nounce")
	}
	//state data has user identifier and the auth provider
	//find the oauth config data by the auth provider and user
	authEnv, err := configProvider.GetOauthConfigOfUser(stateTokenData.GetProvider())
	if err != nil {
		return fmt.Errorf("failed to get auth config of provider %v", err)
	}
	//create
	cfg, err := ConfigFromJSON(authEnv.Secret, stateTokenData.GetScope())

	if err != nil {
		return fmt.Errorf("could not get the oauth config of the provider %v", err)
	}

	if cfg == nil {
		return fmt.Errorf("no secret found matching the name")
	}

	token, err := cfg.Exchange(context.TODO(), authcode)
	if err != nil {
		return fmt.Errorf("failed to exchange with oauth provider to get access token from auth code %v", err)
	}

	ts := authEnv.OauthTokenStorage
	if err != nil {
		return fmt.Errorf("error creating token storage %v", err)
	}
	err = ts.SaveNewToken(token)
	if err != nil {
		return fmt.Errorf("failed to save auth token %v", err)
	}
	return nil
}

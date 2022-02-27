package gotoauth

//UserOauthConfig wraps the oatuh secret (for example the oauth2 config JSON) and the token storage specfic to the user
type UserOauthConfig struct {
	Secret            []byte
	OauthTokenStorage TokenStorage
}

//OAuthConfigProvider defines the functions that the config provider (for example AWS, Local or database) implement to provide the needed configuration data
type OAuthConfigProvider interface {
	//GetOauthConfigOfUser function returns the oauth config data of the given user with the given oauth provider (Google, Atlanssian etc)
	GetOauthConfigOfUser(oauthProvider string) (*UserOauthConfig, error)
}

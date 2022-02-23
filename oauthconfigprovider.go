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

//OauthState defines the functions to manage the data saved with the auth state(nounce), it is expected to have the provide data and user data in it.
type OauthState interface {
	GetStateData() []byte
	GetProvider() string
	GetScope() []string
}

type OauthNounceStateWriter interface {
	Save(nounce string, state OauthState) error
}

type OauthNounceStateReader interface {
	Read(nounce string) (OauthState, error)
}

/*
func NewHttpClient(oauthEnv OauthEnv, scope []string) (*http.Client, error) {

	oauthcfg, err := ConfigFromJSON(oauthEnv.SecretJson, scope)
	if err != nil {
		return nil, err
	}
	tokenStorage := oauthEnv.TokenStorage
	return NewClient(tokenStorage, oauthcfg)
}
*/
/*
const (
	GOOGLE     = "GOOGLE"
	ATLANSSIAN = "ATLANSSIAN"
	GITHUB     = "GITHUB"
)
*/
//type OauthProviders map[string]OauthEnv

/*
var ProviderScope = map[string]string{
	//GOOGLE:     "googlecalendar.CalendarReadonlyScope",
	GOOGLE:     "https://www.googleapis.com/auth/calendar.events.readonly https://www.googleapis.com/auth/gmail.send",
	ATLANSSIAN: "offline_access read:jira-user read:jira-work write:jira-work read:confluence-user write:confluence-content read:confluence-content.all read:confluence-space.summary",
	GITHUB:     "repo",
}
*/
//type GetOauthProviders func() (OauthProviders, error)

//GetAuthEnvOfProvider call the given GetOauthProviders function and return the OauthEnv of the specified provider
/*
func GetAuthEnvOfProvider(env OauthConfigProvider, authProvider string) (*OauthEnv, error) {
	authProviders, err := env.GetOauthProviders()
	if err != nil {
		return nil, err
	}
	authEnv := (authProviders)[authProvider]
	if authEnv.SecretJson == nil {
		return nil, fmt.Errorf("missing auth env for %s", authProvider)
	}
	return &authEnv, nil

}
*/

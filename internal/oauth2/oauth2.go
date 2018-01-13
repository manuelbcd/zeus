package oauth2

import (
	"github.com/google/uuid"
	"github.com/marco2704/zeus/internal/config"
	"github.com/marco2704/zeus/internal/users"
	"github.com/marco2704/zeus/pkg/redis"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"net/http"
	"os"
)

const (
	// stateTimeout is the time-out in seconds given to SetWithExpire call in newState.
	stateTimeout = 120
)

// Client is the interface that unifies oauth2 provider client struct.
type Client interface {
	GetUser(ctx context.Context, code string) (*users.User, error)
	AuthCodeURL() string
}

// oauth2Client unexported base oauth2 client for all client providers.
type oauth2Client struct {
	config *oauth2.Config
}

// getNewClient exchanges the given code for a token, which is used for getting a http.Client.
// The http.Client and a nil error is returned if all went well.
func (client *oauth2Client) getNewClient(ctx context.Context, code string) (*http.Client, error) {

	token, err := client.config.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: token.AccessToken,
		},
	)

	return oauth2.NewClient(ctx, ts), nil
}

func GetClientByState(state string) Client {

	unused, err := checkState(state)
	if err != nil || !unused {
		return nil
	}

	// TODO: un-code provider
	return GetClient(state)
}

//
func GetClient(provider string) Client {

	switch provider {
	case "GITHUB":
		return gitHubClient
	}

	return nil
}

// init initializes an auth.Config struct.
func init() {
	gitHubClient = &GitHubClient{
		&oauth2Client{
			config: &oauth2.Config{
				ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
				ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
				Scopes:       []string{},

				Endpoint: github.Endpoint,
			},
		},
	}
}

// newState generates a random uuid, then this is set with stateTimeout. if all
// were well the uuid will be returned as string, otherwise a empty with a no nil
// error will be returned.
func newState() (string, error) {

	state, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	conn, err := redis.NewSession(config.Config.RedisAddress())
	if err != nil {
		return "", err
	}

	err = redis.SetWithExpire(
		conn,
		state.String(),
		true,
		stateTimeout,
		true,
	)

	if err != nil {
		return "", err
	}

	return state.String(), err
}

// checkState gets the value using state param as key. If the value is true,
// false is set as new value for the next time this state is checked.
func checkState(state string) (bool, error) {

	conn, err := redis.NewSession(config.Config.RedisAddress())
	if err != nil {
		return false, err
	}

	value, err := redis.GetBoolValue(conn, state, false)
	if err != nil {
		return false, err
	}

	if value {
		return value, redis.Set(conn, state, false, true)
	}

	return value, nil
}

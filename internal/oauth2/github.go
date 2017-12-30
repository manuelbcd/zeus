package oauth2

import (
	"github.com/google/go-github/github"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	constants "golang.org/x/oauth2/github"
	"os"
)

var (
	conf *oauth2.Config
)

// init initializes a auth.Config for using GitHub oauth2.
func init() {
	conf = &oauth2.Config{

		ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		Scopes:       []string{},

		Endpoint: constants.Endpoint,
	}
}

// GetUser using given code.
func GetUser(ctx context.Context, code string) (*github.User, error) {

	token, err := conf.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token.AccessToken},
	)

	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	user, _, err := client.Users.Get(ctx, "")
	return user, err
}

// ProviderAuthURL returns a URL to OAuth 2.0 provider's consent page that asks for permissions
// for the required scopes explicitly.
func ProviderAuthURL(state string) string {
	return conf.AuthCodeURL(state)
}

package oauth2

import (
	"github.com/google/go-github/github"
	"github.com/marco2704/zeus/internal/users"
	"golang.org/x/net/context"
	"log"
)

var (
	gitHubClient *GitHubClient
)

// GitHubClient represents a GitHub oauth2Client. It implements client interface.
type GitHubClient struct {
	*oauth2Client
}

// AuthCodeURL is a layer over AuthCodeURL config method creating a new random
// state which is used as state parameter.
// TODO: generate state by provider name
func (client *GitHubClient) AuthCodeURL() string {

	state, err := newState()
	if err != nil {
		log.Println(err)
	}

	return client.config.AuthCodeURL(state)
}

// GetUser gets a GitHub user using code parameter, then create and returns a User
// by calling users.CreateUser method.
func (client *GitHubClient) GetUser(ctx context.Context, code string) (*users.User, error) {

	tc, err := gitHubClient.getNewClient(ctx, code)
	if err != nil {
		return nil, err
	}

	newClient := github.NewClient(tc)
	user, _, err := newClient.Users.Get(ctx, "")

	if err != nil {
		return nil, err
	}

	return users.CreateUser(user.GetEmail(), user.GetName())
}

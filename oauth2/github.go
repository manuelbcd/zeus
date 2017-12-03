//
package oauth2

import (
    "os"
    "log"
    "golang.org/x/oauth2"
    "golang.org/x/net/context"
    oauth2Github"golang.org/x/oauth2/github"
    googleGithub "github.com/google/go-github/github"
)

var (

    conf = &oauth2.Config{

        ClientID: os.Getenv("GITHUB_CLIENT_ID"),
        ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
        Scopes: []string{

        },

        Endpoint: oauth2Github.Endpoint,
    }
)

//
func GetUser(ctx context.Context, code string) (*googleGithub.User, error){

    token, err := conf.Exchange(ctx, code)
    if err != nil {
        return nil, err
    }

    ts := oauth2.StaticTokenSource(
        &oauth2.Token{ AccessToken: token.AccessToken },
    )

    tc := oauth2.NewClient(ctx, ts)
    client := googleGithub.NewClient(tc)
    user, _ , err  := client.Users.Get(ctx,"")
    return user, err
}


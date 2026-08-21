package cupra

import (
	"context"

	"github.com/evcc-io/evcc/plugin/auth"
	"github.com/evcc-io/evcc/vehicle/vag/cariad"
	"golang.org/x/oauth2"
)

// ClientID is the MyCupra Android app client ID (idkClientIDAndroidLive).
const ClientID = "3c756d46-f1ba-4d78-9f9a-cff0d5292d51@apps_vw-dilab_com"

func init() {
	auth.Register("cupra", func(_ map[string]any) (oauth2.TokenSource, error) {
		return NewOAuth(context.Background(), "")
	})
}

// OAuthConfig returns the OAuth2 config for the MyCupra device authorization flow.
func OAuthConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID: ClientID,
		Endpoint: oauth2.Endpoint{
			DeviceAuthURL: "https://identity.vwgroup.io/oidc/v1/device_authorization",
			TokenURL:      cariad.BaseURL + "/auth/v1/idk/oidc/token",
			AuthStyle:     oauth2.AuthStyleInParams,
		},
		Scopes: []string{"openid", "mbb", "profile"},
	}
}

// NewOAuth creates a token source for the MyCupra device authorization flow.
func NewOAuth(ctx context.Context, title string) (oauth2.TokenSource, error) {
	return auth.NewOAuth(ctx, "MyCupra", title, OAuthConfig(), auth.WithOauthDeviceFlowOption())
}

package audi

import (
	"context"

	"github.com/evcc-io/evcc/plugin/auth"
	"golang.org/x/oauth2"
)

// ClientID is the myAudi Android app client id (idkClientIDAndroidLive).
// Retrieved from https://content.app.my.audi.com/service/mobileapp/configurations/market/DE/de?v=4.23.1
const ClientID = "09b6cbec-cd19-4589-82fd-363dfa8c24da@apps_vw-dilab_com"

func init() {
	auth.Register("audi", func(_ map[string]any) (oauth2.TokenSource, error) {
		return NewOAuth(context.Background(), "")
	})
}

func OAuthConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID: ClientID,
		Endpoint: oauth2.Endpoint{
			// Device Authorization Grant (RFC 8628) — does not require Play Integrity attestation
			DeviceAuthURL: "https://identity.vwgroup.io/oidc/v1/device_authorization",
			TokenURL:      "https://emea.bff.cariad.digital/auth/v1/idk/oidc/token",
			AuthStyle:     oauth2.AuthStyleInParams,
		},
		Scopes: []string{"openid", "mbb", "profile"},
	}
}

// NewOAuth creates a new OAuth token source using the Device Authorization Grant.
// On first use the evcc GUI will display a user_code and verification_uri for the
// user to authorise via browser. Subsequent restarts use the stored refresh_token.
func NewOAuth(ctx context.Context, title string) (oauth2.TokenSource, error) {
	return auth.NewOAuth(ctx, "Audi", title, OAuthConfig(), auth.WithOauthDeviceFlowOption())
}

package vehicle

import (
	"context"
	"strings"
	"time"

	"github.com/evcc-io/evcc/api"
	"github.com/evcc-io/evcc/util"
	"github.com/evcc-io/evcc/util/request"
	_ "github.com/evcc-io/evcc/vehicle/audi" // register service endpoint and auth plugin
	"github.com/evcc-io/evcc/vehicle/audi"
	"github.com/evcc-io/evcc/vehicle/vw/weconnect"
)

// Audi is an api.Vehicle implementation for Audi cars using the myAudi API
type Audi struct {
	*embed
	*weconnect.Provider
}

func init() {
	registry.AddCtx("audi", NewAudiFromConfig)
}

// NewAudiFromConfig creates a new Audi vehicle
func NewAudiFromConfig(ctx context.Context, other map[string]any) (api.Vehicle, error) {
	cc := struct {
		embed            `mapstructure:",squash"`
		User, Password   string // deprecated: ignored, kept for config compatibility
		VIN              string
		Cache            time.Duration
		Timeout          time.Duration
	}{
		Cache:   interval,
		Timeout: request.Timeout,
	}

	if err := util.DecodeOther(other, &cc); err != nil {
		return nil, err
	}

	log := util.NewLogger("audi").Redact(cc.VIN)
	authCtx := util.WithLogger(context.Background(), log)

	ts, err := audi.NewOAuth(authCtx, cc.embed.GetTitle())
	if err != nil {
		return nil, err
	}

	wcAPI := weconnect.NewAPI(log, ts)
	wcAPI.Client.Timeout = cc.Timeout

	v := &Audi{
		embed:    &cc.embed,
		Provider: weconnect.NewProvider(wcAPI, strings.ToUpper(cc.VIN), cc.Cache),
	}

	return v, nil
}

package audi

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"

	"github.com/evcc-io/evcc/server/service"
	"github.com/evcc-io/evcc/util"
	"github.com/evcc-io/evcc/vehicle/vw/weconnect"
)

func init() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /vehicles", getVehicles)
	service.Register("audi", mux)
}

// getVehicles lists VINs linked to the myAudi account.
// Reuses the shared OAuth instance so results appear once the user has authorised via the GUI.
func getVehicles(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vins := []string{}
	defer func() { _ = json.NewEncoder(w).Encode(vins) }()

	log := util.NewLogger("audi")
	ctx := util.WithLogger(context.Background(), log)

	ts, err := NewOAuth(ctx, "")
	if err != nil {
		return
	}

	vehicles, err := weconnect.NewAPI(log, ts).Vehicles()
	if err != nil {
		return
	}

	for _, v := range vehicles {
		vins = append(vins, v.VIN)
	}
	slices.Sort(vins)
}

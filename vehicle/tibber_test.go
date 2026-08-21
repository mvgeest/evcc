package vehicle

import (
	"testing"

	"github.com/evcc-io/evcc/api"
	"github.com/evcc-io/evcc/vehicle/tibber"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func makeDeviceDetail(connector, charging string) tibber.DeviceDetail {
	caps := []tibber.Capability{}
	if connector != "" {
		caps = append(caps, tibber.Capability{ID: "connector.status", Value: connector})
	}
	if charging != "" {
		caps = append(caps, tibber.Capability{ID: "charging.status", Value: charging})
	}
	return tibber.DeviceDetail{Capabilities: caps}
}

func TestTibberStatus(t *testing.T) {
	tc := []struct {
		name      string
		connector string // "" = capability absent
		charging  string // "" = capability absent
		want      api.ChargeStatus
	}{
		// Explicit states: connector.status is the primary signal.
		{"disconnected/idle", tibber.StatusDisconnected, tibber.StatusIdle, api.StatusA},
		{"disconnected/unknown", tibber.StatusDisconnected, "unknown", api.StatusA},
		{"connected/idle", tibber.StatusConnected, tibber.StatusIdle, api.StatusB},
		{"connected/unknown", tibber.StatusConnected, "unknown", api.StatusB},
		{"connected/charging", tibber.StatusConnected, tibber.StatusCharging, api.StatusC},

		// charging.status beats connector.status when actively charging.
		{"disconnected/charging", tibber.StatusDisconnected, tibber.StatusCharging, api.StatusC},
		{"unknown-connector/charging", "unknown", tibber.StatusCharging, api.StatusC},

		// OEM reports connector.status as "unknown": fall back to charging.status.
		{"unknown-connector/unknown-charging", "unknown", "unknown", api.StatusA},
		{"unknown-connector/idle", "unknown", tibber.StatusIdle, api.StatusB},
		{"unknown-connector/charging", "unknown", tibber.StatusCharging, api.StatusC},

		// Capability absent.
		{"absent connector, absent charging", "", "", api.StatusA},
		{"absent connector, idle", "", tibber.StatusIdle, api.StatusB},
		{"absent connector, charging", "", tibber.StatusCharging, api.StatusC},
	}

	for _, tc := range tc {
		t.Run(tc.name, func(t *testing.T) {
			detail := makeDeviceDetail(tc.connector, tc.charging)
			v := &Tibber{
				embed: &embed{},
				dataG: func() (tibber.DeviceDetail, error) { return detail, nil },
			}
			status, err := v.Status()
			require.NoError(t, err)
			assert.Equal(t, tc.want, status)
		})
	}
}

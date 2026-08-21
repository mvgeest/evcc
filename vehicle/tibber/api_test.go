package tibber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func makeDetail(connectorStatus, chargingStatus string) DeviceDetail {
	caps := []Capability{}
	if connectorStatus != "" {
		caps = append(caps, Capability{ID: idConnector, Value: connectorStatus})
	}
	if chargingStatus != "" {
		caps = append(caps, Capability{ID: idCharging, Value: chargingStatus})
	}
	return DeviceDetail{Capabilities: caps}
}

func TestPlugStatus(t *testing.T) {
	cases := []struct {
		connector string
		wantVal   string
		wantOk    bool
	}{
		{StatusConnected, StatusConnected, true},
		{StatusDisconnected, StatusDisconnected, true},
		{"unknown", "unknown", true},
		{"", "", false}, // capability absent
	}
	for _, tc := range cases {
		d := makeDetail(tc.connector, "")
		val, ok := d.PlugStatus()
		assert.Equal(t, tc.wantOk, ok, "connector=%q", tc.connector)
		if ok {
			assert.Equal(t, tc.wantVal, val, "connector=%q", tc.connector)
		}
	}
}

func TestChargingStatus(t *testing.T) {
	cases := []struct {
		charging string
		wantVal  string
		wantOk   bool
	}{
		{StatusCharging, StatusCharging, true},
		{StatusIdle, StatusIdle, true},
		{"unknown", "unknown", true},
		{"", "", false},
	}
	for _, tc := range cases {
		d := makeDetail("", tc.charging)
		val, ok := d.ChargingStatus()
		assert.Equal(t, tc.wantOk, ok, "charging=%q", tc.charging)
		if ok {
			assert.Equal(t, tc.wantVal, val, "charging=%q", tc.charging)
		}
	}
}

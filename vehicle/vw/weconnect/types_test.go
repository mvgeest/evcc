package weconnect

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// sampleStatus reflects the actual selectivestatus API response shape.
const sampleStatus = `{
	"charging": {
		"batteryStatus": {
			"value": {
				"currentSOC_pct": 46,
				"cruisingRangeElectric_km": 303
			}
		},
		"chargingStatus": {
			"value": {
				"chargingState": "readyForCharging",
				"remainingChargingTimeToComplete_min": 90
			}
		},
		"plugStatus": {
			"value": {
				"plugConnectionState": "connected"
			}
		},
		"chargingSettings": {
			"value": {
				"targetSOC_pct": 80
			}
		}
	},
	"measurements": {
		"odometerStatus": {
			"value": {
				"odometer": 12345.6
			}
		}
	},
	"climatisation": {
		"climatisationStatus": {
			"value": {
				"climatisationState": "cooling"
			}
		}
	}
}`

func TestStatusParsing(t *testing.T) {
	var s Status
	require.NoError(t, json.Unmarshal([]byte(sampleStatus), &s))

	require.NotNil(t, s.Charging, "charging block must be present")
	assert.Equal(t, 46, s.Charging.BatteryStatus.Value.CurrentSOCPct)
	assert.Equal(t, 303, s.Charging.BatteryStatus.Value.CruisingRangeElectricKm)
	assert.Equal(t, "readyForCharging", s.Charging.ChargingStatus.Value.ChargingState)
	assert.Equal(t, 90, s.Charging.ChargingStatus.Value.RemainingChargingTimeToCompleteMin)
	assert.Equal(t, "connected", s.Charging.PlugStatus.Value.PlugConnectionState)
	require.NotNil(t, s.Charging.ChargingSettings.Value.TargetSOCPct)
	assert.Equal(t, 80, *s.Charging.ChargingSettings.Value.TargetSOCPct)

	require.NotNil(t, s.Measurements)
	assert.Equal(t, 12345.6, s.Measurements.OdometerStatus.Value.Odometer)

	require.NotNil(t, s.Climatisation)
	assert.Equal(t, "cooling", s.Climatisation.ClimatisationStatus.Value.ClimatisationState)
}

func TestParkingPositionLatLon(t *testing.T) {
	raw := `{"lat": 52.123, "lon": 4.567}`
	var p ParkingPosition
	require.NoError(t, json.Unmarshal([]byte(raw), &p))
	assert.Equal(t, 52.123, p.Latitude)
	assert.Equal(t, 4.567, p.Longitude)
}

func TestParkingPositionLatitudeLongitude(t *testing.T) {
	raw := `{"latitude": 52.123, "longitude": 4.567}`
	var p ParkingPosition
	require.NoError(t, json.Unmarshal([]byte(raw), &p))
	assert.Equal(t, 52.123, p.Latitude)
	assert.Equal(t, 4.567, p.Longitude)
}

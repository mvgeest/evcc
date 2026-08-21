package weconnect

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

// Vehicles is the /vehicles api
type Vehicles struct {
	Data []Vehicle
}

// Vehicle is the api vehicle
type Vehicle struct {
	VIN      string
	Model    string
	Nickname string
}

// Status is the /status api
type Status struct {
	Access *struct {
		AccessStatus struct {
			Value struct {
				OverallStatus        string    `json:"overallStatus"`
				CarCapturedTimestamp Timestamp `json:"carCapturedTimestamp"`
				Doors                []struct {
					Name   string   `json:"name"`
					Status []string `json:"status"`
				} `json:"doors"`
				Windows []struct {
					Name   string   `json:"name"`
					Status []string `json:"status"`
				} `json:"windows"`
				DoorLockStatus string `json:"doorLockStatus"`
			} `json:"value"`
		} `json:"accessStatus"`
	} `json:"access"`
	Automation *struct {
		ClimatisationTimer struct {
			Value struct {
				Timers []struct {
					ID             int  `json:"id"`
					Enabled        bool `json:"enabled"`
					RecurringTimer struct {
						StartTime   string `json:"startTime"`
						RecurringOn struct {
							Mondays    bool `json:"mondays"`
							Tuesdays   bool `json:"tuesdays"`
							Wednesdays bool `json:"wednesdays"`
							Thursdays  bool `json:"thursdays"`
							Fridays    bool `json:"fridays"`
							Saturdays  bool `json:"saturdays"`
							Sundays    bool `json:"sundays"`
						} `json:"recurringOn"`
					} `json:"recurringTimer,omitempty"`
					SingleTimer struct {
						StartDateTime Timestamp `json:"startDateTime"`
					} `json:"singleTimer,omitempty"`
				} `json:"timers"`
				CarCapturedTimestamp Timestamp `json:"carCapturedTimestamp"`
				TimeInCar            string    `json:"timeInCar"`
			} `json:"value"`
		} `json:"climatisationTimer"`
		ChargingProfiles *struct {
			Value struct {
				CarCapturedTimestamp Timestamp `json:"carCapturedTimestamp"`
				TimeInCar            string    `json:"timeInCar"`
				Profiles             []any     `json:"profiles"`
			} `json:"value"`
		} `json:"chargingProfiles"`
	} `json:"automation"`
	UserCapabilities *struct {
		CapabilitiesStatus struct {
			Value []struct {
				ID                   string    `json:"id"`
				UserDisablingAllowed bool      `json:"userDisablingAllowed"`
				ExpirationDate       Timestamp `json:"expirationDate,omitempty"`
				Status               []int     `json:"status,omitempty"`
			} `json:"value"`
		} `json:"capabilitiesStatus"`
	} `json:"userCapabilities"`
	Charging *struct {
		BatteryStatus struct {
			Value struct {
				CarCapturedTimestamp    Timestamp `json:"carCapturedTimestamp"`
				CurrentSOCPct           int       `json:"currentSOC_pct"`
				CruisingRangeElectricKm int       `json:"cruisingRangeElectric_km"`
			} `json:"value"`
		} `json:"batteryStatus"`
		ChargingStatus struct {
			Value struct {
				CarCapturedTimestamp               Timestamp `json:"carCapturedTimestamp"`
				RemainingChargingTimeToCompleteMin int       `json:"remainingChargingTimeToComplete_min"`
				ChargingState                      string    `json:"chargingState"` // readyForCharging/off/charging
				ChargeMode                         string    `json:"chargeMode"`
				ChargePowerKW                      float64   `json:"chargePower_kW"`
				ChargeRateLomh                     float64   `json:"chargeRate_lomh"`
				ChargeType                         string    `json:"chargeType"`
				ChargingSettings                   string    `json:"chargingSettings"`
			} `json:"value"`
		} `json:"chargingStatus"`
		ChargingSettings struct {
			Value struct {
				CarCapturedTimestamp Timestamp `json:"carCapturedTimestamp"`
				MaxChargeCurrentAC   string    `json:"maxChargeCurrentAC"`
				AutoUnlockPlugWhenCharged string `json:"autoUnlockPlugWhenCharged"`
				AutoUnlockPlugWhenChargedAC string `json:"autoUnlockPlugWhenCharged_AC"`
				TargetSOCPct         *int      `json:"targetSOC_pct"`
			} `json:"value"`
		} `json:"chargingSettings"`
		PlugStatus struct {
			Value struct {
				CarCapturedTimestamp Timestamp `json:"carCapturedTimestamp"`
				PlugConnectionState  string    `json:"plugConnectionState"` // connected/disconnected
				PlugLockState        string    `json:"plugLockState"`
				ExternalPower        string    `json:"externalPower"`
				LedColor             string    `json:"ledColor"`
			} `json:"value"`
		} `json:"plugStatus"`
		ChargingProfiles *struct {
			Value struct {
				CarCapturedTimestamp Timestamp `json:"carCapturedTimestamp"`
				TimeInCar            string    `json:"timeInCar"`
				Profiles             []struct {
					ProfileID       int    `json:"profileID"`
					ProfileName     string `json:"profileName"`
					ProfileActive   bool   `json:"profileActive"`
					ChargingOptions struct {
						TargetSOCPct        int    `json:"targetSOC_pct"`
						MinSOCPct           int    `json:"minSOC_pct"`
						AutoUnlockPlug      bool   `json:"autoUnlockPlug"`
						AutoUnlockPlugAC    bool   `json:"autoUnlockPlug_AC"`
						PreferredChargeTimes []any `json:"preferredChargeTimes"`
					} `json:"chargingOptions"`
					TimerConstraints struct {
						Timers        []any `json:"timers"`
						MinSOCEnabled bool  `json:"minSOC_enabled"`
					} `json:"timerConstraints"`
				} `json:"profiles"`
			} `json:"value"`
		} `json:"chargingProfilesStatus"`
	} `json:"charging"`
	Climatisation *struct {
		ClimatisationStatus struct {
			Value struct {
				CarCapturedTimestamp  Timestamp `json:"carCapturedTimestamp"`
				RemainingClimatisationTimeMin int `json:"remainingClimatisationTime_min"`
				ClimatisationState   string    `json:"climatisationState"`
			} `json:"value"`
		} `json:"climatisationStatus"`
		ClimatisationSettings struct {
			Value struct {
				CarCapturedTimestamp       Timestamp `json:"carCapturedTimestamp"`
				TargetTemperatureC         float64   `json:"targetTemperature_C"`
				TargetTemperatureF         float64   `json:"targetTemperature_F"`
				UnitInCar                  string    `json:"unitInCar"`
				ClimatisationWithoutExternalPower bool `json:"climatisationWithoutExternalPower"`
			} `json:"value"`
		} `json:"climatisationSettings"`
	} `json:"climatisation"`
	Measurements *struct {
		OdometerStatus struct {
			Value struct {
				CarCapturedTimestamp Timestamp `json:"carCapturedTimestamp"`
				Odometer             float64   `json:"odometer"`
			} `json:"value"`
		} `json:"odometerStatus"`
		FuelLevelStatus struct {
			Value struct {
				CarCapturedTimestamp Timestamp `json:"carCapturedTimestamp"`
				CurrentFuelLevel_pct int       `json:"currentFuelLevel_pct"`
				CurrentSOC_pct       int       `json:"currentSOC_pct"`
				PrimaryEngineType    string    `json:"primaryEngineType"`
				CarType              string    `json:"carType"`
			} `json:"value"`
		} `json:"fuelLevelStatus"`
	} `json:"measurements"`
}

// FuelStatus is the engine range status
type FuelStatus struct {
	RangeStatus struct {
		Value struct {
			CarCapturedTimestamp Timestamp         `json:"carCapturedTimestamp"`
			CarType              string            `json:"carType"`
			PrimaryEngine        EngineRangeStatus `json:"primaryEngine"`
			SecondaryEngine      EngineRangeStatus `json:"secondaryEngine"`
			TotalRangeKm         int               `json:"totalRange_km"`
		} `json:"value"`
	} `json:"rangeStatus"`
}

func (f *FuelStatus) EngineRangeStatus(typ string) (EngineRangeStatus, error) {
	if f == nil {
		return EngineRangeStatus{}, errors.New("missing fuel status")
	}

	if f.RangeStatus.Value.PrimaryEngine.Type == typ {
		return f.RangeStatus.Value.PrimaryEngine, nil
	}
	if f.RangeStatus.Value.SecondaryEngine.Type == typ {
		return f.RangeStatus.Value.SecondaryEngine, nil
	}

	return EngineRangeStatus{}, fmt.Errorf("unknown engine type: %s, got [%s, %s]", typ, f.RangeStatus.Value.PrimaryEngine.Type, f.RangeStatus.Value.SecondaryEngine.Type)
}

// EngineRangeStatus is the engine range status
type EngineRangeStatus struct {
	Type             string `json:"type"`
	CurrentSOCPct    int    `json:"currentSOC_pct"`
	RemainingRangeKm int    `json:"remainingRange_km"`
}

// ParkingPosition is the /parkingposition api response
type ParkingPosition struct {
	Latitude             float64
	Longitude            float64
	CarCapturedTimestamp Timestamp
}

// UnmarshalJSON accepts both the latitude/longitude and the shorter lat/lon
// key variants the parkingposition endpoint may return.
func (p *ParkingPosition) UnmarshalJSON(data []byte) error {
	var res struct {
		Latitude             *float64  `json:"latitude"`
		Longitude            *float64  `json:"longitude"`
		Lat                  *float64  `json:"lat"`
		Lon                  *float64  `json:"lon"`
		CarCapturedTimestamp Timestamp `json:"carCapturedTimestamp"`
	}

	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	p.CarCapturedTimestamp = res.CarCapturedTimestamp

	switch {
	case res.Latitude != nil:
		p.Latitude = *res.Latitude
	case res.Lat != nil:
		p.Latitude = *res.Lat
	}

	switch {
	case res.Longitude != nil:
		p.Longitude = *res.Longitude
	case res.Lon != nil:
		p.Longitude = *res.Lon
	}

	return nil
}

// Timestamp implements JSON unmarshal
type Timestamp struct {
	time.Time
}

// UnmarshalJSON decodes string timestamp into time.Time
func (ct *Timestamp) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), "\"")

	t, err := time.Parse(time.RFC3339, s)
	if err == nil {
		ct.Time = t
	}

	return err
}

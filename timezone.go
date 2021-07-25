package go_android_utils

import (
	"encoding/json"
	"fmt"
	"time"
)

type TimeZone struct {
	Location     *time.Location
	LocationName string // Location doesn't get marshalled, use this for persistence
}

type auxTimeZone struct {
	Location string `json:"location"`
}

func (tz *TimeZone) FromName(name string) error {
	loc, err := time.LoadLocation(name)
	if err == nil {
		tz.LocationName = name
		tz.Location = loc
		return err
	}
	return fmt.Errorf("TimeZone.FromName: %w", err)
}

func (tz *TimeZone) FromLocation(loc *time.Location) {
	tz.LocationName = loc.String()
	tz.Location = loc
}

func (tz *TimeZone) GetLocation() *time.Location {
	result := tz.Location
	return result
}

func (tz *TimeZone) GetLocationName() string {
	result := tz.LocationName
	return result
}

func NewTimeZoneFromName(name string) (*TimeZone, error) {
	tz := &TimeZone{}
	err := tz.FromName(name)
	return tz, err
}

func NewTimeZoneFromLocation(location *time.Location) *TimeZone {
	tz := &TimeZone{}
	tz.FromLocation(location)
	return tz
}

func (tz *TimeZone) MarshalJSON() ([]byte, error) {
	return json.Marshal(&auxTimeZone{
		Location: tz.GetLocationName(),
	})
}

func (tz *TimeZone) UnmarshalJSON(data []byte) error {
	aux := &auxTimeZone{}
	err := json.Unmarshal(data, aux)
	if err == nil {
		err = tz.FromName(aux.Location)
	}
	return err
}

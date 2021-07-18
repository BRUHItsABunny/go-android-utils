package go_android_utils

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

type TimeZone struct {
	sync.RWMutex
	Location     *time.Location
	LocationName string // Location doesn't get marshalled, use this for persistence
}

type auxTimeZone struct {
	Location string `json:"location"`
}

func (tz *TimeZone) FromName(name string) error {
	loc, err := time.LoadLocation(name)
	if err == nil {
		tz.Lock()
		tz.LocationName = name
		tz.Location = loc
		tz.Unlock()
		return err
	}
	return fmt.Errorf("TimeZone.FromName: %w", err)
}

func (tz *TimeZone) FromLocation(loc *time.Location) {
	tz.Lock()
	tz.LocationName = loc.String()
	tz.Location = loc
	tz.Unlock()
}

func (tz *TimeZone) GetLocation() *time.Location {
	tz.RLock()
	result := tz.Location
	tz.RUnlock()
	return result
}

func (tz *TimeZone) GetLocationName() string {
	tz.RLock()
	result := tz.LocationName
	tz.RUnlock()
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

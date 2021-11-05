package go_android_utils

import (
	"fmt"
	"time"
)

func (tz *Timezone) FromName(name string) error {
	_, err := time.LoadLocation(name)
	if err == nil {
		tz.Name = name
		return err
	}
	return fmt.Errorf("TimeZone.FromName: %w", err)
}

func (tz *Timezone) FromLocation(loc *time.Location) {
	tz.Name = loc.String()
}

func (tz *Timezone) MustGoLocation() *time.Location {
	result, err := time.LoadLocation(tz.Name)
	if err != nil {
		return time.UTC
	}
	return result
}

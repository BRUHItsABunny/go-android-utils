package go_android_utils

import "strings"

func (location *GPSLocation) Accuracy() int {
	if location.Provider == 0 {
		return randomInt(1, 3)
	}
	return int(location.Provider)
}

func (location *GPSLocation) ProviderString() string {
	provider := location.Provider
	if location.Provider == 0 {
		provider = LocationProvider(randomInt(1, 3))
	}
	return strings.ToLower(LocationProvider_name[int32(provider)])
}

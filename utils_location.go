package go_android_utils

func (location *GPSLocation) Accuracy() int {
	if location.Provider == 0 {
		return randomInt(1, 3)
	}
	return int(location.Provider)
}

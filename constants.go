package go_android_utils

var (
	DefaultAndroidID       = NewAndroidID()
	DefaultDeviceOnePlus7T = Device{
		AndroidId:            NewAndroidID(),
		Locale:               &Locale{Language: "en", Country: "us"},
		AndroidVersion:       &AndroidVersion{sdk: 29},
		Device:               "OnePlus7T",
		Manufacturer:         "OnePlus",
		Model:                "HD1905",
		Product:              "OnePlus7T",
		Build:                "QKQ1.190716.003",
		Type:                 "user",
		Tags:                 "release-keys",
		IncrementalVersion:   "2101212100",
		DPI:                  420,
		ResolutionHorizontal: 1080,
		ResolutionVertical:   2400,
		Architecture:         &Architecture{cpu: "arm64"},
	}
)

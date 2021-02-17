package go_android_utils

import (
	"errors"
	"strconv"
	"strings"
)

type AndroidVersion struct {
	sdk int
}

func (version *AndroidVersion) FromAndroidVersion(versionStr string) error {
	var err error
	// 6 => 6.0, 6.1 => 6.1
	if !strings.Contains(versionStr, ".") {
		versionStr += ".0"
	}
	sdk, ok := AndroidVersionToSDK[versionStr]
	if ok {
		version.sdk, _ = strconv.Atoi(sdk)
	} else {
		err = ErrAndroidVersionVersionUnsupported
	}
	return err
}

func (version *AndroidVersion) FromAndroidSDK(sdkStr string) error {
	var err error
	_, ok := AndroidSDKToVersion[sdkStr]
	if ok {
		sdk, _ := strconv.ParseInt(sdkStr, 10, 64)
		version.sdk = int(sdk)
	} else {
		err = ErrAndroidVersionSDKUnsupported
	}
	return err
}

func (version AndroidVersion) ToAndroidVersion() string {
	versionStr := AndroidSDKToVersion[strconv.Itoa(version.sdk)]
	better, ok := AndroidVersionToAndroidVersion[versionStr]
	if ok {
		return better
	}
	return versionStr
}

func (version AndroidVersion) ToAndroidSDK() string {
	return strconv.Itoa(version.sdk)
}

func (version AndroidVersion) IsOlder(comparison AndroidVersion) bool {
	return version.sdk < comparison.sdk
}

func (version AndroidVersion) IsNewer(comparison AndroidVersion) bool {
	return version.sdk > comparison.sdk
}

func (version AndroidVersion) Equals(comparison AndroidVersion) bool {
	return version.sdk == comparison.sdk
}

// DATABASE

var (
	ErrAndroidVersionSDKUnsupported     = errors.New("the supplied SDK is unsupported")
	ErrAndroidVersionVersionUnsupported = errors.New("the supplied version is unsupported")

	AndroidSDKToVersion = map[string]string{
		"30": "11.0",
		"29": "10.0",
		"28": "9.0",
		"27": "8.1",
		"26": "8.0",
		"25": "7.1",
		"24": "7",
		"23": "6.0",
		"22": "5.1",
		"21": "5.0",
		"19": "4.4",
		"18": "4.3",
		"17": "4.2",
		"16": "4.1",
		"14": "4.0",
	}

	AndroidVersionToSDK = map[string]string{
		"11.0": "30",
		"10.0": "29",
		"9.0":  "28",
		"8.1":  "27",
		"8.0":  "26",
		"7.1":  "25",
		"7":    "24",
		"6.0":  "23",
		"5.1":  "22",
		"5.0":  "21",
		"4.4":  "19",
		"4.3":  "18",
		"4.2":  "17",
		"4.1":  "16",
		"4.0":  "14",
	}
	AndroidVersionToAndroidVersion = map[string]string{
		"11.0": "11",
		"10.0": "10",
		"9.0":  "9",
		"8.0":  "8",
		"7.0":  "7",
		"6.0":  "6",
		"5.0":  "5",
		"4.0":  "4",
	}
)

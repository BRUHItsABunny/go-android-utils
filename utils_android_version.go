package go_android_utils

import (
	"errors"
	"strconv"
	"strings"
)

func (version AndroidVersion) IsValid() bool {
	if version == 0 {
		return false
	}
	_, result := AndroidVersion_name[int32(version)]
	return result
}

func AndroidVersionFromVersionString(versionStr string) (AndroidVersion, error) {
	var err error
	// 6 => 6.0, 6.1 => 6.1
	versionSplits := strings.Split(versionStr, ".")
	if len(versionSplits) == 1 {
		versionSplits = append(versionSplits, "0")
	}
	val, ok := AndroidVersion_value["AndroidVersion_"+strings.ToUpper(strings.Join(versionSplits[:2], "_"))]
	if !ok {
		err = ErrAndroidVersionVersionUnsupported
		val = 0
	}
	return AndroidVersion(val), err
}

func AndroidVersionFromSDKString(sdkStr string) (AndroidVersion, error) {
	sdk, err := strconv.ParseInt(sdkStr, 10, 64)
	if err == nil {
		_, ok := AndroidVersion_name[int32(sdk)]
		if !ok || sdk < 1 {
			err = ErrAndroidVersionSDKUnsupported
		}
		return AndroidVersion(int(sdk)), err
	}
	return AndroidVersion(0), err
}

func (version AndroidVersion) ToAndroidVersion() string {
	return strings.ReplaceAll(version.String()[15:], "_", ".")
}

func (version AndroidVersion) ToAndroidSDK() string {
	return strconv.Itoa(int(version))
}

var (
	ErrAndroidVersionSDKUnsupported     = errors.New("the supplied SDK is unsupported")
	ErrAndroidVersionVersionUnsupported = errors.New("the supplied version is unsupported")
)

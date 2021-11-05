package go_android_utils

import (
	"strconv"
	"strings"
)

func DeviceFromFingerprint(fingerprint string) (*Device, error) {
	// "OnePlus/OnePlus5/OnePlus5:9/PKQ1.180716.001/2002242003:user/release-keys"
	var (
		err    error
		device = new(Device)
	)
	mainParts := strings.Split(fingerprint, "/")
	device.Manufacturer = mainParts[0]
	device.Product = mainParts[1]
	subParts := strings.Split(mainParts[2], ":")
	device.Device = subParts[0]
	device.Version, err = AndroidVersionFromVersionString(subParts[1])
	if err == nil {
		device.Build = mainParts[3]
		subParts = strings.Split(mainParts[4], ":")
		device.IncrementalVersion = subParts[0]
		device.Type = subParts[1]
		device.Tags = mainParts[5]
	}
	return device, err
}

func DeviceFromUserAgent(userAgent string) (*Device, error) {
	// Mozilla/5.0 (Linux; Android 4.2.1; en-us; Nexus 5 Build/JOP40D)
	var (
		err    error
		device = new(Device)
	)
	indexStart := strings.Index(userAgent, "(")
	indexStop := strings.Index(userAgent, ")")
	deviceStr := userAgent[indexStart:indexStop]
	for _, elem := range strings.Split(deviceStr, "; ") {
		if strings.Contains(elem, "Android ") {
			device.Version, err = AndroidVersionFromVersionString(strings.Split(elem, " ")[1])
		} else if strings.Contains(elem, "-") {
			device.Locale, err = LocaleFromLocaleString(elem)
		} else if strings.Contains(elem, "Build/") {
			parts := strings.Split(elem, " Build/")
			device.Model = parts[0]
			device.Build = parts[1]
		}

		if err != nil {
			break
		}
	}

	return device, err
}

func (device *Device) GetUserAgent() string {
	// Returns the device string part of useragent, manually need to prefix and postfix it with data like what software/browser the user agent is supposed to be
	result := "(Linux; Android " + device.Version.ToAndroidVersion() + "; " + strings.ToLower(device.Locale.ToLocale("-", true)) + "; " + device.Model + " Build/" + device.Build + ")"
	return result
}

const (
	DeviceFormatKeyAndroidVersion  = ":andVers"
	DeviceFormatKeyAndroidSDKLevel = ":andSDK"
	DeviceFormatKeyLocale          = ":locale"
	DeviceFormatKeyModel           = ":model"
	DeviceFormatKeyBuild           = ":build"
	DeviceFormatKeyDPI             = ":dpi"
	DeviceFormatKeyDevice          = ":device"
	DeviceFormatKeyManufacturer    = ":manufacturer"
)

func (device *Device) FormatUserAgent(format string) string {
	// TODO: Cache this? replace is inefficient everytime?
	format = strings.ReplaceAll(format, DeviceFormatKeyAndroidVersion, device.Version.ToAndroidVersion())
	format = strings.ReplaceAll(format, DeviceFormatKeyAndroidSDKLevel, device.Version.ToAndroidSDK())
	format = strings.ReplaceAll(format, DeviceFormatKeyLocale, strings.ToLower(device.Locale.ToLocale("-", true)))
	format = strings.ReplaceAll(format, DeviceFormatKeyModel, device.Model)
	format = strings.ReplaceAll(format, DeviceFormatKeyBuild, device.Build)
	format = strings.ReplaceAll(format, DeviceFormatKeyDPI, strconv.Itoa(int(device.DPI)))
	format = strings.ReplaceAll(format, DeviceFormatKeyDevice, device.Device)
	format = strings.ReplaceAll(format, DeviceFormatKeyManufacturer, device.Manufacturer)
	return format
}

func (device *Device) GetFingerprint() string {
	result := device.Manufacturer + "/" + device.Product + "/" + device.Device + ":" + device.Version.ToAndroidVersion() + "/" + device.Build + "/" + device.IncrementalVersion + ":" + device.Type + "/" + device.Tags
	return result
}

package go_android_utils

import (
	"encoding/json"
	"strconv"
	"strings"
)

type Device struct {
	AndroidId            *AndroidID      // random
	Locale               *Locale         // en-us
	AndroidVersion       *AndroidVersion // 9 (translates to sdk 28)
	Device               string          // OnePlus5
	Manufacturer         string          // OnePlus
	Model                string          // ONEPLUS A5000
	Product              string          // OnePlus5
	Build                string          // PKQ1.180716.001
	Type                 string          // user
	Tags                 string          // release-keys
	IncrementalVersion   string          // 2002242003
	DPI                  int             //
	ResolutionHorizontal int             //
	ResolutionVertical   int             //
	Architecture         *Architecture   // ARM64
	TimeZone             *TimeZone       // America/Chicago
}

type auxDevice struct {
	AndroidId            *AndroidID      `json:"android_id"`
	Locale               *Locale         `json:"locale"`
	AndroidVersion       *AndroidVersion `json:"android_version"`
	Device               string          `json:"device"`
	Manufacturer         string          `json:"manufacturer"`
	Model                string          `json:"model"`
	Product              string          `json:"product"`
	Build                string          `json:"build"`
	Type                 string          `json:"type"`
	Tags                 string          `json:"tags"`
	IncrementalVersion   string          `json:"rom_version"`
	DPI                  int             `json:"dpi"`
	ResolutionHorizontal int             `json:"resolution_horizontal"`
	ResolutionVertical   int             `json:"resolution_vertical"`
	Architecture         *Architecture   `json:"architecture"`
	TimeZone             *TimeZone       `json:"time_zone"`
}

func (device *Device) FromFingerprint(fingerprint string) error {
	// "OnePlus/OnePlus5/OnePlus5:9/PKQ1.180716.001/2002242003:user/release-keys"
	var err error
	mainParts := strings.Split(fingerprint, "/")
	device.Manufacturer = mainParts[0]
	device.Product = mainParts[1]
	subParts := strings.Split(mainParts[2], ":")
	device.Device = subParts[0]
	device.AndroidVersion = &AndroidVersion{}
	err = device.AndroidVersion.FromAndroidVersion(subParts[1])
	if err == nil {
		device.Build = mainParts[3]
		subParts = strings.Split(mainParts[4], ":")
		device.IncrementalVersion = subParts[0]
		device.Type = subParts[1]
		device.Tags = mainParts[5]
	}
	return err
}

func (device *Device) FromUserAgent(userAgent string) error {
	// Mozilla/5.0 (Linux; Android 4.2.1; en-us; Nexus 5 Build/JOP40D)
	var err error
	indexStart := strings.Index(userAgent, "(")
	indexStop := strings.Index(userAgent, ")")
	deviceStr := userAgent[indexStart:indexStop]
	for _, elem := range strings.Split(deviceStr, "; ") {
		if strings.Contains(elem, "Android ") {
			androidVersion := AndroidVersion{}
			err = androidVersion.FromAndroidVersion(strings.Split(elem, " ")[1])
			if err == nil {
				device.AndroidVersion = &androidVersion
			}
		} else if strings.Contains(elem, "-") {
			locale := Locale{}
			err = locale.FromLocale(elem)
			if err == nil {
				device.Locale = &locale
			}
		} else if strings.Contains(elem, "Build/") {
			parts := strings.Split(elem, " Build/")
			device.Model = parts[0]
			device.Build = parts[1]
		}

		if err != nil {
			break
		}
	}

	return err
}

func (device *Device) GetUserAgent() string {
	// Returns the device string part of useragent, manually need to prefix and postfix it with data like what software/browser the user agent is supposed to be
	result := "(Linux; Android " + device.AndroidVersion.ToAndroidVersion() + "; " + strings.ToLower(device.Locale.ToLocale("-")) + "; " + device.Model + " Build/" + device.Build + ")"
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
	format = strings.ReplaceAll(format, DeviceFormatKeyAndroidVersion, device.AndroidVersion.ToAndroidVersion())
	format = strings.ReplaceAll(format, DeviceFormatKeyAndroidSDKLevel, device.AndroidVersion.ToAndroidSDK())
	format = strings.ReplaceAll(format, DeviceFormatKeyLocale, strings.ToLower(device.Locale.ToLocale("-")))
	format = strings.ReplaceAll(format, DeviceFormatKeyModel, device.Model)
	format = strings.ReplaceAll(format, DeviceFormatKeyBuild, device.Build)
	format = strings.ReplaceAll(format, DeviceFormatKeyDPI, strconv.Itoa(device.DPI))
	format = strings.ReplaceAll(format, DeviceFormatKeyDevice, device.Device)
	format = strings.ReplaceAll(format, DeviceFormatKeyManufacturer, device.Manufacturer)
	return format
}

func (device *Device) GetFingerprint() string {
	result := device.Manufacturer + "/" + device.Product + "/" + device.Device + ":" + device.AndroidVersion.ToAndroidVersion() + "/" + device.Build + "/" + device.IncrementalVersion + ":" + device.Type + "/" + device.Tags
	return result
}

func (device *Device) MarshalJSON() ([]byte, error) {

	aux := &auxDevice{
		AndroidId:            device.AndroidId,
		Locale:               device.Locale,
		AndroidVersion:       device.AndroidVersion,
		Device:               device.Device,
		Manufacturer:         device.Manufacturer,
		Model:                device.Model,
		Product:              device.Product,
		Build:                device.Build,
		Type:                 device.Type,
		Tags:                 device.Tags,
		IncrementalVersion:   device.IncrementalVersion,
		DPI:                  device.DPI,
		ResolutionHorizontal: device.ResolutionHorizontal,
		ResolutionVertical:   device.ResolutionVertical,
		Architecture:         device.Architecture,
		TimeZone:             device.TimeZone,
	}

	return json.Marshal(aux)
}

func (device *Device) UnmarshalJSON(data []byte) error {
	aux := &auxDevice{}
	err := json.Unmarshal(data, aux)
	if err == nil {
		device.AndroidId = aux.AndroidId
		device.Locale = aux.Locale
		device.AndroidVersion = aux.AndroidVersion
		device.Device = aux.Device
		device.Manufacturer = aux.Manufacturer
		device.Model = aux.Model
		device.Product = aux.Product
		device.Build = aux.Build
		device.Type = aux.Type
		device.Tags = aux.Tags
		device.IncrementalVersion = aux.IncrementalVersion
		device.DPI = aux.DPI
		device.ResolutionHorizontal = aux.ResolutionHorizontal
		device.ResolutionVertical = aux.ResolutionVertical
		device.Architecture = aux.Architecture
		device.TimeZone = aux.TimeZone
	}
	return err
}

func (device *Device) GetAndroidID() *AndroidID {
	result := device.AndroidId
	return result
}

func (device *Device) GetLocale() *Locale {
	result := device.Locale
	return result
}

func (device *Device) GetAndroidVersion() *AndroidVersion {
	result := device.AndroidVersion
	return result
}

func (device *Device) GetArchitecture() *Architecture {
	result := device.Architecture
	return result
}

func (device *Device) GetTimeZone() *TimeZone {
	result := device.TimeZone
	return result
}

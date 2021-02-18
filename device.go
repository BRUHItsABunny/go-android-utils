package go_android_utils

import "strings"

type Device struct {
	AndroidId            AndroidID
	Locale               Locale         // en-us
	AndroidVersion       AndroidVersion // 9 (translates to sdk 28)
	Device               string         // OnePlus5
	Manufacturer         string         // OnePlus
	Model                string         // ONEPLUS A5000
	Product              string         // OnePlus5
	Build                string         // PKQ1.180716.001
	Type                 string         // user
	Tags                 string         // release-keys
	IncrementalVersion   string         // 2002242003
	DPI                  int
	ResolutionHorizontal int
	ResolutionVertical   int
	ABIS                 []string
}

func (device *Device) FromFingerprint(fingerprint string) error {
	// dev=OnePlus5, man=OnePlus, mod="ONEPLUS A5000", pro=OnePlus5, fin="OnePlus/OnePlus5/OnePlus5:9/PKQ1.180716.001/2002242003:user/release-keys", sdk=28
	var err error
	fingerprint = strings.ReplaceAll(fingerprint, "\"", "")
	mainParts := strings.Split(fingerprint, ", ")
	for _, mainPart := range mainParts {
		subParts := strings.Split(mainPart, "=")
		switch subParts[0] {
		case "dev":
			device.Device = subParts[1]
			break
		case "man":
			device.Manufacturer = subParts[1]
			break
		case "mod":
			device.Model = subParts[1]
			break
		case "pro":
			device.Product = subParts[1]
			break
		case "fin":
			// "OnePlus/OnePlus5/OnePlus5:9/PKQ1.180716.001/2002242003:user/release-keys"
			fingerprintParts := strings.Split(subParts[1], "/")
			device.Build = fingerprintParts[3]
			OTAParts := strings.Split(fingerprintParts[4], ":")
			device.IncrementalVersion = OTAParts[0]
			device.Type = OTAParts[1]
			device.Tags = fingerprintParts[5]
			break
		case "sdk":
			androidVersion := AndroidVersion{}
			err = androidVersion.FromAndroidSDK(subParts[1])
			if err == nil {
				device.AndroidVersion = androidVersion
			}
			break
		}
		if err != nil {
			break
		}
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
				device.AndroidVersion = androidVersion
			}
		} else if strings.Contains(elem, "-") {
			locale := Locale{}
			err = locale.FromLocale(elem)
			if err == nil {
				device.Locale = locale
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

func (device Device) GetUserAgent() string {
	// Returns the device string part of useragent, manually need to prefix and postfix it with data like what software/browser the user agent is supposed to be
	return "(Linux; Android " + device.AndroidVersion.ToAndroidVersion() + "; " + strings.ToLower(device.Locale.ToLocale("-")) + "; " + device.Model + " Build/" + device.Build + ")"
}

func (device Device) GetFingerprint() string {
	return "dev=" + device.Device + ", man=" + device.Manufacturer + ", mod=\"" + device.Model + "\", pro=" + device.Product + ", fin=\"" + device.Manufacturer + "/" + device.Product + "/" + device.Device + ":" + device.AndroidVersion.ToAndroidVersion() + "/" + device.Build + "/" + device.IncrementalVersion + ":" + device.Type + "/" + device.Tags + "\", sdk=" + device.AndroidVersion.ToAndroidSDK()
}

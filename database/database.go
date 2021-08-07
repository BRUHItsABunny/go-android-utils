package database

import (
	"encoding/json"
	go_android_utils "github.com/BRUHItsABunny/go-android-utils"
	"math/rand"
	"time"
)

// Here we store a few devices and way to get them, just easy access in case you want to prototype a few devices in a library fast
// TODO: Actually add a few devices here...
var DeviceDB = map[string]string{
	// "oneplus3": "",
	"oneplus5":  "{\"android_id\":{\"id\":\"163f5f0f9a621d72\"},\"locale\":{\"language\":\"en\",\"country\":\"us\"},\"android_version\":{\"sdk\":28},\"device\":\"OnePlus5\",\"manufacturer\":\"OnePlus\",\"model\":\"ONEPLUS A5000\",\"product\":\"OnePlus5\",\"build\":\"PKQ1.180716.001\",\"type\":\"user\",\"tags\":\"release-keys\",\"rom_version\":\"2002242003\",\"dpi\":420,\"resolution_horizontal\":1080,\"resolution_vertical\":1920,\"architecture\":{\"architecture\":\"arm64\"},\"time_zone\":{\"location\":\"UTC\"}}",
	"oneplus7t": "{\"android_id\":{\"id\":\"163f5f0f9a621d72\"},\"locale\":{\"language\":\"en\",\"country\":\"us\"},\"android_version\":{\"sdk\":29},\"device\":\"OnePlus7T\",\"manufacturer\":\"OnePlus\",\"model\":\"HD1905\",\"product\":\"OnePlus7T\",\"build\":\"QKQ1.190716.003\",\"type\":\"user\",\"tags\":\"release-keys\",\"rom_version\":\"2101212100\",\"dpi\":420,\"resolution_horizontal\":1080,\"resolution_vertical\":2400,\"architecture\":{\"architecture\":\"arm64\"},\"time_zone\":{\"location\":\"UTC\"}}",
	// "oneplus9": "",
	"oneplus9pro": "{\"android_id\":{\"id\":\"163f5f0f9a621d72\"},\"locale\":{\"language\":\"en\",\"country\":\"us\"},\"android_version\":{\"sdk\":30},\"device\":\"OnePlus9Pro\",\"manufacturer\":\"OnePlus\",\"model\":\"LE2125\",\"product\":\"OnePlus9Pro\",\"build\":\"RKQ1.201105.002\",\"type\":\"user\",\"tags\":\"release-keys\",\"rom_version\":\"2105290043\",\"dpi\":600,\"resolution_horizontal\":1440,\"resolution_vertical\":3216,\"architecture\":{\"architecture\":\"arm64\"},\"time_zone\":{\"location\":\"UTC\"}}",
}

var DeviceDBKeys = []string{
	// "oneplus3",
	"oneplus5",
	"oneplus7t",
	// "oneplus9",
	"oneplus9pro",
}

func GetDBDevice(key string) (*go_android_utils.Device, bool) {
	device := new(go_android_utils.Device)
	val, found := DeviceDB[key]
	if found {
		_ = json.Unmarshal([]byte(val), device)
	}
	// Device from DB needs to be random ID
	_ = device.AndroidId.Random()
	return device, found
}

func GetRandomDevice() *go_android_utils.Device {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	device := new(go_android_utils.Device)
	val := DeviceDB[DeviceDBKeys[r.Intn(len(DeviceDBKeys))]]
	_ = json.Unmarshal([]byte(val), device)
	// Device from DB needs to be random ID
	_ = device.AndroidId.Random()
	return device
}

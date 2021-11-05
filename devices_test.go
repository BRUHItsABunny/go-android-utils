package go_android_utils

import (
	"fmt"
	"testing"
	"time"
)

func TestDeviceFromFingerprint(t *testing.T) {
	fingerPrint := "OnePlus/OnePlus5/OnePlus5:9/PKQ1.180716.001/2002242003:user/release-keys"
	device, err := DeviceFromFingerprint(fingerPrint)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(device.String())
}

func TestDeviceFromUserAgent(t *testing.T) {
	userAgent := "Mozilla/5.0 (Linux; Android 4.2.1; en-us; Nexus 5 Build/JOP40D)"
	device, err := DeviceFromUserAgent(userAgent)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(device.String())
	fmt.Println(device)
}

func TestRandomDevices(t *testing.T) {
	for i := 0; i < 10; i++ {
		device := GetRandomDevice()
		fmt.Println(device.String())
		fmt.Println(device.Location.ProviderString())
		time.Sleep(5 * time.Second)
	}
}

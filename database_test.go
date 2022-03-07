package go_android_utils

import (
	"context"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func TestDBIFaceJSON(t *testing.T) {
	location := "dump.json" // You may need your own dump.json

	db, err := NewJSONAUDDevice(location)
	if err != nil {
		t.Error()
	}

	device, err := db.GetDevice(context.Background())
	if err != nil {
		t.Error(err)
	}

	device.Locale = &Locale{
		Language:   "nl",
		CountryISO: "NL",
	}
	device.Randomize()

	fmt.Println(fmt.Sprintf("%#v", device))
	fmt.Println(spew.Sdump(device))
}

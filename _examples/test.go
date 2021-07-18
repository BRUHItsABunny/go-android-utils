package main

import (
	"encoding/json"
	"fmt"
	go_android_utils "github.com/BRUHItsABunny/go-android-utils"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	device := new(go_android_utils.Device)
	err := json.Unmarshal([]byte("{\"android_id\":{\"id\":\"163f5f0f9a621d72\"},\"locale\":{\"language\":\"en\",\"country\":\"us\"},\"android_version\":{\"sdk\":29},\"device\":\"OnePlus7T\",\"manufacturer\":\"OnePlus\",\"model\":\"HD1905\",\"product\":\"OnePlus7T\",\"build\":\"QKQ1.190716.003\",\"type\":\"user\",\"tags\":\"release-keys\",\"rom_version\":\"2101212100\",\"dpi\":420,\"resolution_horizontal\":1080,\"resolution_vertical\":2400,\"architecture\":{\"architecture\":\"arm64\"},\"time_zone\":{\"location\":\"UTC\"}}"), device)
	if err == nil {
		fmt.Println(spew.Sdump(device))
	} else {
		fmt.Println("err: ", err)
	}
}

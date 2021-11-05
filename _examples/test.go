package main

import (
	"fmt"
	go_android_utils "github.com/BRUHItsABunny/go-android-utils"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	device := new(go_android_utils.Device)
	device = go_android_utils.GetRandomDevice()
	fmt.Println(spew.Sdump(device))
}

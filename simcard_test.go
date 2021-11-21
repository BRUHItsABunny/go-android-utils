package go_android_utils

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

// Extract from .txt
func TestBunny(t *testing.T) {
	f, _ := os.Open("mnc.txt")
	body, _ := io.ReadAll(f)
	bodyStr := string(body)

	countries := map[string][]string{}
	lines := strings.Split(bodyStr, "\n")
	for _, line := range lines {
		components := strings.Split(line, "\t")
		country := strings.ToUpper(components[2])
		entries, exists := countries[country]
		if !exists {
			entries = []string{}
		}

		entries = append(entries, "{MNC: \""+components[1]+"\", MCC: \""+components[0]+"\"},")
		countries[country] = entries
	}

	/*
		"GP": {
				{MNC: "08", MCC: "340"},
				{MNC: "10", MCC: "340"},
			},
	*/
	for k, v := range countries {
		fmt.Println("\"" + k + "\": {")
		for _, e := range v {
			fmt.Println("    " + e)
		}
		fmt.Println("},")
	}
}

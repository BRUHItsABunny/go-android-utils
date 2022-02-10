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
	f, _ := os.Open("./_resources/samples/mnc.txt")
	body, _ := io.ReadAll(f)
	bodyStr := string(body)

	countries := map[string][]string{}
	lines := strings.Split(bodyStr, "\r\n")
	for _, line := range lines {
		// fmt.Println(fmt.Sprintf("%#v", line))
		components := strings.Split(line, "\t")
		country := strings.ToUpper(components[2])
		entries, exists := countries[country]
		if !exists {
			entries = []string{}
		}

		if len(components) == 5 {
			components = append(components, "")
		}
		entries = append(entries, fmt.Sprintf("{MNC: \"%s\", MCC: \"%s\", Carrier: \"%s\", CountryISO: \"%s\", CountryCode: \"%s\"},", components[1], components[0], components[5], strings.ToUpper(components[2]), components[4]))
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

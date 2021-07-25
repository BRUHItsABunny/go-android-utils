package go_android_utils

import (
	"encoding/json"
	"errors"
)

type Architecture struct {
	cpu string
}

type auxArchitecture struct {
	Architecture string `json:"architecture"`
}

func (architecture *Architecture) FromArchitecture(architectureStr string) error {
	_, ok := strInSlice(AvailableArchitectures, architectureStr)
	if ok {
		architecture.cpu = architectureStr
		return nil
	}
	return ErrArchitectureNotSupported
}

func (architecture *Architecture) ToArchitecture() string {
	result := architecture.cpu
	return result
}

func (architecture *Architecture) ToABI() []string {
	result := architecture.cpu
	return DefaultArchitectures[result]
}

func (architecture *Architecture) MarshalJSON() ([]byte, error) {
	return json.Marshal(&auxArchitecture{
		Architecture: architecture.ToArchitecture(),
	})
}

func (architecture *Architecture) UnmarshalJSON(data []byte) error {
	aux := &auxArchitecture{}
	err := json.Unmarshal(data, aux)
	if err == nil {
		err = architecture.FromArchitecture(aux.Architecture)
	}
	return err
}

var (
	ErrArchitectureNotSupported = errors.New("the supplied CPU architecture is unsupported")

	DefaultArchitectures = map[string][]string{
		"armeabi": []string{"armeabi-v7a", "armeabi"},
		"arm64":   []string{"arm64-v8a", "arm64"},
		"x86":     []string{"x86", "x86"},
		"x86_64":  []string{"x86_64", "x86_64"},
	}
	AvailableArchitectures = []string{
		"armeabi",
		"arm64",
		"x86",
		"x86_64",
	}
)

package go_android_utils

import "errors"

type Architecture struct {
	cpu string
}

func (architecture *Architecture) FromArchitecture(architectureStr string) error {
	_, ok := strInSlice(AvailableArchitectures, architectureStr)
	if ok {
		architecture.cpu = architectureStr
	}
	return ErrArchitectureNotSupported
}

func (architecture Architecture) ToArchitecture() string {
	return architecture.cpu
}

func (architecture Architecture) ToABI() []string {
	return DefaultArchitectures[architecture.cpu]
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

package go_android_utils

import "errors"

func (arch Architecture) IsValid() bool {
	if arch == 0 {
		return false
	}
	_, result := Architecture_name[int32(arch)]
	return result
}

func (arch Architecture) ToABI() ([]string, error) {
	val, ok := DefaultABIS[arch]
	if !ok {
		return nil, ErrArchitectureNotSupported
	} else {
		return val, nil
	}
}

var (
	DefaultABIS = map[Architecture][]string{
		Architecture_Architecture_ARMEABI: []string{"armeabi-v7a", "armeabi"},
		Architecture_Architecture_ARM64:   []string{"arm64-v8a", "arm64"},
		Architecture_Architecture_X86:     []string{"x86", "x86"},
		Architecture_Architecture_X86_64:  []string{"x86_64", "x86_64"},
	}

	ErrArchitectureNotSupported = errors.New("the supplied CPU architecture is unsupported")
)

package go_android_utils

import "strings"

func isNumeric(c rune) bool { return c < '0' || c > '9' }

func IsNumeric(in string) bool {
	return strings.IndexFunc(in, isNumeric) == -1
}

func (s *SIMCard) IsValid() bool {
	return IsNumeric(s.MNC) && IsNumeric(s.MCC) && IsNumeric(s.PhoneNumber)
}

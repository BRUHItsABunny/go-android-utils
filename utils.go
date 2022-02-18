package go_android_utils

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	// Guarantee a seed?
	rand.Seed(time.Now().UnixNano())
}

func stringsAreNotNull(elems ...string) (int, bool) {
	for i, str := range elems {
		if str == "" {
			return i, false
		}
	}
	return -1, true
}

func strInSlice(haystack []string, needle string) (int, bool) {
	for i, elem := range haystack {
		if elem == needle {
			return i, true
		}
	}
	return -1, false
}

func randomInt(min, max int) int {
	return rand.Intn(max-min) + min
}

func randomStrSlice(strSlice []string) string {
	return strSlice[rand.Intn(len(strSlice))]
}

func randomSIMSlice(strSlice []*SIMCard) *SIMCard {
	return strSlice[rand.Intn(len(strSlice))]
}

func removeAllNONHex(r rune) rune {
	if r > 47 && r < 58 {
		return r
	} else if r > 96 && r < 103 {
		return r
	}
	return -1
}

func groupSubString(in, fill string, length int) []string {
	result := make([]string, 0)
	inLen := len(in)
	for i := 0; i < inLen; {
		j := i + length
		if j >= inLen {
			j = inLen
		}
		group := in[i:j]
		for len(group) < length {
			group += fill
		}
		result = append(result, group[0:length])
		i = j
	}
	return result
}

func isNumeric(c rune) bool { return c < '0' || c > '9' }

func IsNumeric(in string) bool {
	return strings.IndexFunc(in, isNumeric) == -1
}

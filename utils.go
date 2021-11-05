package go_android_utils

import (
	"math/rand"
	"time"
)

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
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(max-min) + min
}

func randomStrSlice(strSlice []string) string {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return strSlice[r.Intn(len(strSlice))]
}

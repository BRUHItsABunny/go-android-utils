package go_android_utils

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

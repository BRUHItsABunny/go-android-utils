package go_android_utils

func stringsAreNotNull(elems ...string) (int, bool) {
	for i, str := range elems {
		if str == "" {
			return i, false
		}
	}
	return -1, true
}

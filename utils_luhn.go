package go_android_utils

// Props to: https://github.com/theplant/luhn
// TODO: Alter to allow for MEID as well, which can be HEXADECIMAL

func LuhnCalculate(number int64) int64 {
	checkNumber := LuhnChecksum(number)

	if checkNumber == 0 {
		return 0
	}
	return 10 - checkNumber
}

func LuhnValid(number int64) bool {
	return (number%10+LuhnChecksum(number/10))%10 == 0
}

func LuhnChecksum(number int64) int64 {
	var luhn int64

	for i := 0; number > 0; i++ {
		cur := number % 10

		if i%2 == 0 {
			cur = cur * 2
			if cur > 9 {
				cur = cur%10 + cur/10
			}
		}

		luhn += cur
		number = number / 10
	}
	return luhn % 10
}

package go_android_utils

import "strings"

func isNumeric(c rune) bool { return c < '0' || c > '9' }

func IsNumeric(in string) bool {
	return strings.IndexFunc(in, isNumeric) == -1
}

func (s *SIMCard) IsValid() bool {
	return IsNumeric(s.MNC) && IsNumeric(s.MCC) && IsNumeric(s.PhoneNumber)
}

// GetHNI emulates the following:
// TelephonyManager tel = (TelephonyManager) getSystemService(Context.TELEPHONY_SERVICE);
// String networkOperator = tel.getNetworkOperator();
func (s *SIMCard) GetHNI() string {
	if IsNumeric(s.MNC) && IsNumeric(s.MCC) {
		return s.MCC + s.MNC
	}
	return "000000"
}

// GetCarrierName Android system returns the names slightly different, there is no way to do that properly for the list of carriers we have now
// heuristic = true at least tries, but it's not 100% accurate
func (s *SIMCard) GetCarrierName(heuristic bool) string {
	if heuristic {
		return strings.Split(s.Carrier, " ")[0]
	}
	return s.Carrier
}

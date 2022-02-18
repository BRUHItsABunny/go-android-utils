package go_android_utils

import (
	"math/rand"
	"strconv"
	"strings"
)

func (s *SIMCard) IsValid() bool {
	return IsNumeric(s.MNC) && IsNumeric(s.MCC) && IsNumeric(s.PhoneNumber)
}

// GetHNI emulates the following:
// TelephonyManager tel = (TelephonyManager) getSystemService(Context.TELEPHONY_SERVICE);
// String networkOperator = tel.getNetworkOperator();
func (s *SIMCard) GetHNI() string {
	if IsNumeric(s.MNC) && IsNumeric(s.MCC) {
		result := s.MCC + s.MNC
		if len(result) < 5 {
			return "000000"
		}
		return result
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

func (s *SIMCard) Randomize(countryISO string) {
	_, ok := AvailableSIMCards[countryISO]
	if !ok {
		countryISO = randomStrSlice(AvailableCountries)
	}
	simCard := randomSIMSlice(AvailableSIMCards[countryISO])
	s.MNC = simCard.MNC
	s.MCC = simCard.MCC
	s.Carrier = simCard.Carrier
	s.CountryCode = simCard.CountryCode
	s.CountryISO = simCard.CountryISO
	if s.Imei == nil {
		s.Imei = new(IMEI)
	}
}

func (i *IMEI) Generate(tac, serial string) (string, error) {
	if len(tac) < 1 {
		tac = i.TAC
	}
	for len(tac) < 8 {
		tac += strconv.Itoa(rand.Intn(9-0) + 0)
	}
	for len(serial) < 6 {
		serial += strconv.Itoa(rand.Intn(9-0) + 0)
	}
	imei := tac + serial
	imeiInt, err := strconv.ParseInt(imei, 10, 64)
	if err == nil {
		imei += strconv.FormatInt(LuhnCalculate(imeiInt), 10)
		i.Imei = imei
	}
	return i.Imei, err
}

/*
func (i *MEID) Generate(region, manuCode, serial string) (string, error) {

}
*/

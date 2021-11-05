package go_android_utils

import (
	"errors"
	"strings"
)

var (
	ErrLocaleFormatUnsupported = errors.New("the supplied locale had an unsupported format")
	ErrLocaleUnsupported       = errors.New("the supplied locale is unsupported")
)

func (locale *Locale) ToLocale(separator string, iso bool) string {
	result := locale.Language + separator + locale.GetCountry(iso)
	return result
}

func (locale *Locale) GetCountry(iso bool) string {
	if iso {
		return strings.ToUpper(locale.CountryISO)
	} else {
		return strings.ToLower(locale.CountryISO)
	}
}

func fromSlice(parts []string) (*Locale, error) {
	result := new(Locale)
	if len(parts) == 2 {
		_, ok := stringsAreNotNull(parts...)
		if ok {
			result.Language = strings.ToLower(parts[0])
			result.CountryISO = strings.ToUpper(parts[1])
		} else {
			return result, ErrLocaleFormatUnsupported
		}
	} else {
		return result, ErrLocaleFormatUnsupported
	}
	return result, nil
}

func LocaleFromLocaleString(localeStr string) (*Locale, error) {
	// Support en_US and en-US
	var err error

	if strings.Contains(localeStr, "-") {
		parts := strings.Split(localeStr, "-")
		return fromSlice(parts)
	} else if strings.Contains(localeStr, "_") {
		parts := strings.Split(localeStr, "_")
		return fromSlice(parts)
	} else {
		err = ErrLocaleFormatUnsupported
	}

	return nil, err
}

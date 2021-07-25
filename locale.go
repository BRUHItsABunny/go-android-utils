package go_android_utils

import (
	"encoding/json"
	"errors"
	"strings"
)

type Locale struct {
	Language string
	Country  string
}

type auxLocale struct {
	Language string `json:"language"`
	Country  string `json:"country"`
}

var (
	ErrLocaleFormatUnsupported = errors.New("the supplied locale had an unsupported format")
	ErrLocaleUnsupported       = errors.New("the supplied locale is unsupported")
)

func (locale *Locale) FromLocale(localeStr string) error {
	// Support en_US and en-US
	var err error

	if strings.Contains(localeStr, "-") {
		parts := strings.Split(localeStr, "-")
		err = locale.fromSlice(parts)
	} else if strings.Contains(localeStr, "_") {
		parts := strings.Split(localeStr, "_")
		err = locale.fromSlice(parts)
	} else {
		err = ErrLocaleFormatUnsupported
	}

	return err
}

func (locale *Locale) FromStrings(language, countryISO string) error {
	var err error

	_, ok := stringsAreNotNull(language, countryISO)
	if ok {
		locale.Language = strings.ToLower(language)
		locale.Country = strings.ToUpper(countryISO)
	} else {
		err = ErrLocaleFormatUnsupported
	}

	return err
}

func (locale *Locale) fromSlice(parts []string) error {
	if len(parts) == 2 {
		_, ok := stringsAreNotNull(parts...)
		if ok {
			locale.Language = strings.ToLower(parts[0])
			locale.Country = strings.ToUpper(parts[1])
		} else {
			return ErrLocaleFormatUnsupported
		}
	} else {
		return ErrLocaleFormatUnsupported
	}
	return nil
}

func (locale *Locale) ToLocale(separator string) string {
	result := locale.Language + separator + locale.Country
	return result
}

func (locale *Locale) GetLanguage() string {
	result := locale.Language
	return result
}

func (locale *Locale) GetCountry() string {
	result := locale.Country
	return result
}

func (locale *Locale) MarshalJSON() ([]byte, error) {

	aux := &auxLocale{
		Country:  locale.Country,
		Language: locale.Language,
	}

	return json.Marshal(aux)
}

func (locale *Locale) UnmarshalJSON(data []byte) error {
	aux := &auxLocale{}
	err := json.Unmarshal(data, aux)
	if err == nil {
		err = locale.FromStrings(aux.Language, aux.Country)
	}
	return err
}

package go_android_utils

import (
	"encoding/hex"
	"math/rand"
	"strconv"
)

type AndroidID struct {
	id uint64
}

func (id *AndroidID) FromHex(idStr string) error {
	result, err := strconv.ParseUint(idStr, 16, 64)
	if err == nil {
		id.id = result
	}
	return err
}

func (id *AndroidID) Random() error {
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err == nil {
		err = id.FromHex(hex.EncodeToString(b))
	}
	return err
}

func (id AndroidID) toDecimalString() string {
	return strconv.FormatUint(id.id, 10)
}

func (id AndroidID) toHexString() string {
	return strconv.FormatUint(id.id, 16)
}

func (id AndroidID) Equals(comparison AndroidID) bool {
	return id.id == comparison.id
}

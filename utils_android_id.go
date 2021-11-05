package go_android_utils

import (
	"encoding/base64"
	"encoding/hex"
	"math/rand"
	"strconv"
)

func NewAndroidID() *AndroidID {
	result := &AndroidID{}
	_ = result.Random()
	return result
}

func (id *AndroidID) FromHex(idStr string) error {
	result, err := strconv.ParseUint(idStr, 16, 64)
	if err == nil {
		id.SetID(result)
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

func (id *AndroidID) ToDecimalString() string {
	return strconv.FormatUint(id.GetID(), 10)
}

func (id *AndroidID) ToHexString() string {
	return strconv.FormatUint(id.GetID(), 16)
}

func (id *AndroidID) ToBase64String() string {
	hByte, _ := hex.DecodeString(id.ToHexString())
	return base64.StdEncoding.EncodeToString(hByte)
}

func (id *AndroidID) Equals(comparison *AndroidID) bool {
	return id.GetID() == comparison.GetID()
}

func (id *AndroidID) IsNull() bool {
	return id.GetID() < 1
}

func (id *AndroidID) GetID() uint64 {
	result := id.Id
	return result
}

func (id *AndroidID) SetID(idN uint64) {
	id.Id = idN
}

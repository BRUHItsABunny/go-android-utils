package go_android_utils

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"math/rand"
	"strconv"
	"sync"
)

type AndroidID struct {
	sync.RWMutex
	id uint64
}

type auxAndroidID struct {
	ID string `json:"id"`
}

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
	id.RLock()
	result := id.id
	id.RUnlock()
	return result
}

func (id *AndroidID) SetID(idN uint64) {
	id.Lock()
	id.id = idN
	id.Unlock()
}

// JSON Custom marshalling, required to even save AID at all

func (id *AndroidID) MarshalJSON() ([]byte, error) {
	return json.Marshal(&auxAndroidID{
		ID: id.ToHexString(),
	})
}

func (id *AndroidID) UnmarshalJSON(data []byte) error {

	aux := &auxAndroidID{}
	err := json.Unmarshal(data, aux)
	if err == nil {
		err = id.FromHex(aux.ID)
	}
	return err
}

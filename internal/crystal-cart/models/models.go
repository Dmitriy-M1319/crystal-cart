package models

import "encoding/json"

type CartProductInfo struct {
	ProductID int64
	Count     int64
}

type Products []CartProductInfo

func (p Products) MarshalBinary() ([]byte, error) {
	return json.Marshal(p)
}

func (p Products) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &p)
}

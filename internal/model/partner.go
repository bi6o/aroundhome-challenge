package model

import (
	"database/sql/driver"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type Partner struct {
	ID                uuid.UUID   `db:"id" json:"id"`
	FlooringMaterials StringSlice `db:"flooring_materials" json:"flooring_materials"`
	AddressLat        float64     `db:"address_lat" json:"address_lat"`
	AddressLong       float64     `db:"address_long" json:"address_long"`
	OperatingRadius   int         `db:"operating_radius" json:"operating_radius"`
	Rating            int         `db:"rating" json:"rating"`
}

type StringSlice []string

func (s StringSlice) Value() (driver.Value, error) {
	if s == nil {
		return nil, nil
	}

	var b strings.Builder

	b.WriteString("{")
	for i, v := range s {
		if i > 0 {
			b.WriteString(",")
		}

		escaped := strings.ReplaceAll(v, "\\", "\\\\")
		escaped = strings.ReplaceAll(escaped, "\"", "\\\"")
		b.WriteString(fmt.Sprintf("\"%s\"", escaped))
	}
	b.WriteString("}")

	return b.String(), nil
}

func (ss *StringSlice) Scan(value interface{}) error {
	byteSlice, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to convert %v to []byte", value)
	}

	str := string(byteSlice)

	str = strings.Trim(str, "{}")
	*ss = strings.Split(str, ",")

	return nil
}

package types

import (
	"errors"
	"fmt"
)

// ConvertibleBoolean is a custom boolean struct used for manipulating
// booleans of string, integer, and boolean to/from JSON and SQL
type ConvertibleBoolean bool

// UnmarshalJSON re-implements the encoding/json Unmarshal method.
// It will cast the any of the truthy/falsy variants to boolean
func (bit *ConvertibleBoolean) UnmarshalJSON(data []byte) error {
	asString := string(data)
	if asString == "1" || asString == "true" {
		*bit = true
	} else if asString == "0" || asString == "false" || asString == "null" {
		*bit = false
	} else {
		return errors.New(fmt.Sprintf("Boolean unmarshal error: invalid input %s", asString))
	}
	return nil
}

// Scan re-implements the database/sql Scan() method.
// It will map the integer found in the response to
// its boolean equivalent
func (bit *ConvertibleBoolean) Scan(src interface{}) error {
	switch src.(type) {
	case int64:
		switch src.(int64) {
		case 0:
			*bit = false
		case 1:
			*bit = true
		default:
			return errors.New("incompatible type for ConvertibleBoolean")
		}
	}
	return nil
}

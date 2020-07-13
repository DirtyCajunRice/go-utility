package types

import (
	"database/sql/driver"
	"encoding/xml"
	"errors"
	"fmt"
	"strconv"
	"time"
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

// UnixTimestamp is a custom time.Time struct used for manipulating
// unix timestamps to/from JSON, XML, and SQL
type UnixTimestamp time.Time

// MarshalJSON re-implements the encoding/json Marshal method.
// It will return a quote-enclosed RFC3339 string as []byte
func (t UnixTimestamp) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", t.String())), nil
}

// UnmarshalXMLAttr re-implements the database/xml Unmarshal()
// method for attributes. It will map the unix timestamp integer
// found in the response to a UnixTimestamp struct
func (t *UnixTimestamp) UnmarshalXMLAttr(attr xml.Attr) error {
	n, err := strconv.ParseInt(attr.Value, 10, 64)
	if err != nil {
		return err
	}
	*t = UnixTimestamp(time.Unix(n, 0))
	return nil
}

// String is a convenience function to get a string
// response based on format. time.RFC3339 is the default
func (t *UnixTimestamp) String(format ...string) string {
	f := time.RFC3339
	if len(format) > 0 {
		f = format[0]
	}
	return time.Time(*t).Format(f)
}

// Scan re-implements the database/sql Scan() method.
// It will convert the RFC 3339 string found in the
// response to a UnixTimestamp struct
func (t *UnixTimestamp) Scan(src interface{}) error {
	switch src.(type) {
	case string:
		if v, err := time.Parse(time.RFC3339, src.(string)); err != nil {
			return err
		} else {
			*t = UnixTimestamp(v)
		}
	}
	return nil
}

// Value re-implements the database/sql Value() method.
// It will convert the custom struct to the standard
// time.Time struct
func (t UnixTimestamp) Value() (driver.Value, error) {
	return time.Time(t), nil
}

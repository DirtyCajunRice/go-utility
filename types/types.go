package types

import (
	"database/sql/driver"
	"encoding/xml"
	"errors"
	"fmt"
	"strconv"
	"time"
)

type ConvertibleBoolean bool

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

type Time time.Time

func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", t.String())), nil
}

func (t *Time) UnmarshalXMLAttr(attr xml.Attr) error {
	n, err := strconv.ParseInt(attr.Value, 10, 64)
	if err != nil {
		return err
	}
	*t = Time(time.Unix(n, 0))
	return nil
}

func (t *Time) String() string {
	return time.Time(*t).Format(time.RFC3339)
}

func (t Time) Value() (driver.Value, error) {
	return time.Time(t), nil
}

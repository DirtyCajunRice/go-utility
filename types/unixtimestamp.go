package types

import (
	"database/sql/driver"
	"encoding/xml"
	"fmt"
	"strconv"
	"time"
)

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
// It will convert the "almost" RFC3339 string found
// in the response to a UnixTimestamp struct
func (t *UnixTimestamp) Scan(src interface{}) error {
	const almostRFC3339 = "2006-01-02 15:04:05-07:00"
	switch src.(type) {
	case string:
		if v, err := time.Parse(almostRFC3339, src.(string)); err != nil {
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

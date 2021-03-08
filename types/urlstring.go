package types

import (
	"database/sql/driver"
	"net/url"
)

// URLString is a custom url.URL struct used for converting
// URLs to/from string for sql
type URLString url.URL

// Scan re-implements the database/sql Scan() method.
// It will convert the url string found
// in the response to a URLString struct
func (u *URLString) Scan(src interface{}) error {

	switch src.(type) {
	case string:
		if v, err := url.Parse(src.(string)); err != nil {
			return err
		} else {
			*u = URLString(*v)
		}
	}
	return nil
}

// Value re-implements the database/sql Value() method.
// It will convert the custom struct to string
func (u URLString) Value() (driver.Value, error) {
	return u.String(), nil
}

// String is a wrapper function of the base url.Url type
func (u *URLString) String() string {
	Url := url.URL(*u)
	return Url.String()
}

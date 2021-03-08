package types

import (
	"fmt"
	"strconv"
)

// IntString is a custom int struct used for manipulating
// ints of string to/from JSON. NOT NEEDED use `json:",string"`
type IntString int

// UnmarshalJSON re-implements the encoding/json Unmarshal method.
// It will cast the string to int type
func (i *IntString) UnmarshalJSON(data []byte) error {
	unquote, err := strconv.Unquote(string(data))
	if err != nil {
		return err
	}
	var integer int
	integer, err = strconv.Atoi(unquote)
	if err != nil {
		return err
	}
	*i = IntString(integer)
	return nil
}

// MarshalJSON re-implements the encoding/json Marshal method.
// It will return a quote-enclosed int as []byte
func (i IntString) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%d\"", i)), nil
}

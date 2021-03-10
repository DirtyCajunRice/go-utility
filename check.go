package utility

import (
	"bytes"
	"reflect"
)

func InStringSlice(slice []string, str string) bool {
	for _, i := range slice {
		if i == str {
			return true
		}
	}
	return false
}

func InIntSlice(slice []int, i int) bool {
	for _, v := range slice {
		if v == i {
			return true
		}
	}
	return false
}

func InSlice(slice interface{}, in interface{}) (bool, error) {
	v := reflect.ValueOf(slice)
	s := make([]interface{}, v.Len())
	for i := 0; i < v.Len(); i++ {
		s[i] = v.Index(i).Interface()
	}
	for _, p := range s {
		if p == in {
			return true, nil
		}
	}
	return false, nil
}

func ArrayOrObject(data []byte) (isArray, isObject bool) {
	x := bytes.TrimLeft(data, " \t\r\n")
	isArray = len(x) > 0 && x[0] == '['
	isObject = len(x) > 0 && x[0] == '{'
	return
}

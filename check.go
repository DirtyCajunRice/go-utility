package utility

import "bytes"

func IsInSlice(slice []string, str string) bool {
	for _, i := range slice {
		if i == str {
			return true
		}
	}
	return false
}

func ArrayOrObject(data []byte) (isArray, isObject bool) {
	x := bytes.TrimLeft(data, " \t\r\n")
	isArray = len(x) > 0 && x[0] == '['
	isObject = len(x) > 0 && x[0] == '{'
	return
}

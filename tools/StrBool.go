package tools

import (
	"strconv"
)

func StrBool(word string) bool {
	boolValue, err := strconv.ParseBool(word)
	if err != nil {
		ZapLogger("console").Error(err.Error())
	}
	return boolValue
}

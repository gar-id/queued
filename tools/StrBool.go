package tools

import (
	"log"
	"strconv"
)

func StrBool(word string) bool {
	boolValue, err := strconv.ParseBool(word)
	if err != nil {
		log.Print(err)
	}
	return boolValue
}

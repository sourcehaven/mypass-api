package util

import (
	"strconv"
)

func ToUint64(val string) uint64 {
	num, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		panic("Value cannot be converted to uint64")
	}
	return num
}

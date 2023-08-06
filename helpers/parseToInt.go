package helpers

import (
	"strconv"
)

func ParseToInteger(data string) (int, error) {
	num, err := strconv.Atoi(data)
	if err != nil {
		return 0, err
	}

	return num, err
}

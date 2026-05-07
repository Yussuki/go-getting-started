package utils

import (
	"errors"
	"strconv"
)

func StringsToFloat(strings []string) ([]float64, error) {
	floats := make([]float64, len(strings))
	for stringIndex, stringValue := range strings {
		floatValue, err := strconv.ParseFloat(stringValue, 64)
		if err != nil {
			return nil, errors.New("Failed to convert string to float.")
		}

		floats[stringIndex] = floatValue
	}

	return floats, nil
}

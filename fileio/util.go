package fileio

import (
	"errors"
)

// ByteArrayToByte asd
func ByteArrayToByte(data []byte) (byte, error) {
	if len(data) != 8 {
		return 0, errors.New("data should have 8 bit length")
	}
	ans := 0
	for _, value := range data {
		ans = ans * 2
		if value < 48 || value > 49 {
			return 0, errors.New("unexpected bit")
		}
		if value == 49 {
			ans = ans + 1
		}
	}
	return (byte)(ans), nil
}

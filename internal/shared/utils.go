package shared

import (
	"encoding/hex"
	"errors"
	"fmt"
	"regexp"
)

func BytesToHex(array []byte) string {
	return hex.EncodeToString(array)
}

func HexToBytes(data string) ([]byte, error) {
	if !IsHex(data) {
		return []byte{}, errors.New("parameter is not hexadecimal")
	} else if len(data) % 2 == 1 {
		return []byte{}, errors.New("parameter is missing last character to be represented in bytes")
	}
	bytes, err := hex.DecodeString(data)
	if err != nil {
		return []byte{}, fmt.Errorf("utils.HexToBytes: %s", err)
	}

	return bytes, nil
}

func HexToBytes16(data string) ([]uint16, error) {
	if len(data)%4 != 0 {
		return []uint16{}, errors.New("parameter is missing last characters to be represented in uint16")
	}

	bytes, err := HexToBytes(data)
	if err != nil {
		return []uint16{}, err
	}

	result := make([]uint16, len(bytes)/2)
	for i := 0; i < len(result); i++ {
		result[i] = uint16(bytes[i*2+1]) | uint16(bytes[i*2])<<8
	}

	return result, nil
}

func StringToBytes(data string) []byte {
	return []byte(data)
}

func HexToBytes32(data string) ([32]byte, error) {
	bytes, err := HexToBytes(data)
	if err != nil {
		return [32]byte{}, fmt.Errorf("utils.HexToBytes32: %s", err)
	}
	var bytes32 [32]byte
	copy(bytes32[:], bytes)

	return bytes32, nil
}

func Stringify(data interface{}) string {
	return fmt.Sprintf("%v", data)
}

func IsHex(h string) bool {
	regexp, _ := regexp.MatchString("^[0-9a-fA-F]+$", h)
	return regexp
}

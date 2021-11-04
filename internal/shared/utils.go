package shared

import (
	"encoding/hex"
	"fmt"
)

func BytesToHex (array []byte) string {
	return hex.EncodeToString(array)
}

func HexToBytes (string string) ([]byte, error) {
	bytes, err := hex.DecodeString(string)
	if err != nil {
		return []byte{}, fmt.Errorf("utils.HexToBytes: %s",err)
	}

	return bytes, nil
}

func StringToBytes (string string) []byte {
	return []byte(string)
}

func HexToBytes32 (string string) ([32]byte, error) {
	bytes, err := HexToBytes(string)
	if err != nil {
		return [32]byte{}, fmt.Errorf("utils.HexToBytes32: %s", err)
	}
	var bytes32 [32]byte
	copy(bytes32[:], bytes)

	return bytes32, nil
}

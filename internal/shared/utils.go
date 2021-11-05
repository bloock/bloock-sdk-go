package shared

import (
	"encoding/hex"
	"fmt"
)

func BytesToHex (array []byte) string {
	return hex.EncodeToString(array)
}

func HexToBytes (data string) ([]byte, error) {
	bytes, err := hex.DecodeString(data)
	if err != nil {
		return []byte{}, fmt.Errorf("utils.HexToBytes: %s",err)
	}

	return bytes, nil
}

func StringToBytes (data string) []byte {
	return []byte(data)
}

func HexToBytes32 (data string) ([32]byte, error) {
	bytes, err := HexToBytes(data)
	if err != nil {
		return [32]byte{}, fmt.Errorf("utils.HexToBytes32: %s", err)
	}
	var bytes32 [32]byte
	copy(bytes32[:], bytes)

	return bytes32, nil
}

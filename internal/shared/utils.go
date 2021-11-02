package shared

import "encoding/hex"

func BytesToHex (array []byte) string {
	return hex.EncodeToString(array)
}

func StringToBytes (string string) []byte {
	return []byte(string)
}

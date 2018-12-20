package hexstring

import (
	"strconv"
	"strings"
)

// HexStringToBytes takes a space separated string of hexadecimal bytes (like "7f")
// and converts them to a slice of bytes.
func HexStringToBytes(hexstring string) ([]byte, error) {
	hexStrings := strings.Split(hexstring, " ")
	hexBytes := make([]byte, 0, len(hexStrings))
	for _, s := range hexStrings {
		if len(s) == 0 {
			continue
		}
		u, err := strconv.ParseUint(s, 16, 8)
		if err != nil {
			return hexBytes, err
		}
		hexBytes = append(hexBytes, uint8(u))
	}
	return hexBytes, nil
}

// BytesToHexString converts a slice of bytes to a space sparated string
// of hexidecimal bytes (like "7f").
func BytesToHexString(bs []byte) string {
	var sb strings.Builder
	for i, b := range bs {
		if i > 0 {
			sb.WriteString(" ")
		}
		sb.WriteString(strconv.FormatUint(uint64(b), 16))
	}
	return sb.String()
}

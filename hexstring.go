package hexstring

import (
	"strconv"
	"strings"
)

func HexStringToBytes(hexstring string) ([]byte, error) {
	hexStrings := strings.Split(hexstring, " ")
	hexBytes := make([]byte, 0, len(hexStrings))
	for _, s := range hexStrings {
		if len(s) == 0 {
			continue
		}
		u, err := strconv.ParseUint(s, 16, 8)
		if err != nil {
			return []byte{}, err
		}
		hexBytes = append(hexBytes, uint8(u))
	}
	return hexBytes, nil
}

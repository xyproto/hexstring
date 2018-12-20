package hexstring

import (
	"fmt"
)

func ExampleHexStringToBytes() {
	helloWorldBytes, err := HexStringToBytes("48 65 6c 6c 6f 20 57 6f 72 6c 64 21")
	if err != nil {
		panic(err)
	}
	fmt.Println(helloWorldBytes)
	// Output: [72 101 108 108 111 32 87 111 114 108 100 33]
}

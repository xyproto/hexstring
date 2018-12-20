package hexstring

import (
	"fmt"
)

func ExampleStringToBytes() {
	helloWorldBytes, err := StringToBytes("48 65 6c 6c 6f 20 57 6f 72 6c 64 21")
	if err != nil {
		panic(err)
	}
	fmt.Println(helloWorldBytes)
	// Output: [72 101 108 108 111 32 87 111 114 108 100 33]
}

func ExampleBytesToString() {
	helloWorldBytes, err := StringToBytes("48 65 6c 6c 6f 20 57 6f 72 6c 64 21")
	if err != nil {
		panic(err)
	}
	fmt.Println(BytesToString(helloWorldBytes))
	// Output: 48 65 6c 6c 6f 20 57 6f 72 6c 64 21
}

func ExampleStringToBytes2() {
	fourBytes, err := StringToBytes("\n01 02\n0a ff\n")
	if err != nil {
		panic(err)
	}
	fmt.Println(fourBytes)
	// Output: [1 2 10 255]
}

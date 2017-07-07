package binary_test

import "fmt"

func Example_addBinaryDigits() {
	data := []bool{false, true, true, false, true, true, false, false}
	var result int
	for _, v := range data {
		result = result << 1
		if v {
			result = result | 1
		}
		fmt.Printf("%d %b\n", result, result)
	}
	// Output:
	// 0 0
	// 1 1
	// 3 11
	// 6 110
	// 13 1101
	// 27 11011
	// 54 110110
	// 108 1101100
}

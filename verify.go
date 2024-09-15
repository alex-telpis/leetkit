package leetkit

import (
	"fmt"
)

// Verify checks if the result is equal to the expected value and prints the outcome.
// Use this function to validate your solution against test cases.
func Verify(expect any, result any) {
	res, exp := Sprint(result), Sprint(expect)
	if res == exp {
		fmt.Printf("    ✅ Pass. TestCase: %s\n", exp)
		return
	}

	fmt.Printf("    ❌ Fail. Expected: %s, but got: %s\n", exp, res)
}

package leetkit

import (
	"fmt"
)

// Verify checks if the result is equal to the expected value and prints the outcome.
// Use this function to validate your solution against test cases.
func Verify(expect any, result any) {
	exp, res := prepareVerify(expect, result)
	if res == exp {
		fmt.Printf("    ✅ Pass. TestCase: %s\n", exp)
		return
	}

	fmt.Printf("    ❌ Fail. Expected: %s, but got: %s\n", exp, res)
}

func prepareVerify(expect any, result any) (string, string) {
	if exp, isStr := expect.(string); isStr {
		if res, isStr := result.(string); isStr {
			return exp, res
		}

		return exp, Sprint(result)
	}

	return Sprint(expect), Sprint(result)
}

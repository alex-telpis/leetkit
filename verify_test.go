package leetkit

import (
	"testing"

	"github.com/alex-telpis/leetkit/internal/testutil"
)

func TestVerify(t *testing.T) {
	cases := []struct {
		expect any
		result any
		equal  bool
	}{
		{expect: 1, result: 1, equal: true},
		{expect: 1, result: "1", equal: false},
		{expect: "abc", result: "abc", equal: true},

		{
			expect: &ListNode{Val: 1, Next: &ListNode{Val: 2}},
			result: &ListNode{Val: 1, Next: &ListNode{Val: 2}},
			equal:  true,
		},
		{
			expect: &ListNode{Val: 1, Next: &ListNode{Val: 2}},
			result: &ListNode{Val: 2, Next: &ListNode{Val: 2}},
			equal:  false,
		},
		{
			expect: "[1,2]",
			result: &ListNode{Val: 1, Next: &ListNode{Val: 2}},
			equal:  true,
		},
		{
			expect: `[2,3]`,
			result: &ListNode{Val: 2, Next: &ListNode{Val: 3}},
			equal:  true,
		},

		{
			expect: `["3","4"]`,
			result: []string{"3", "4"},
			equal:  true,
		},
		{
			expect: `["a3","a4"]`,
			result: []string{"a3", "a4"},
			equal:  true,
		},
		{
			expect: `"["5","6"]"`,
			result: []string{"5", "6"},
			equal:  false,
		},
		{
			expect: `["6","7"]`,
			result: []int{1, 2},
			equal:  false,
		},
	}

	for _, tc := range cases {
		wantStr, gotStr := prepareVerify(tc.expect, tc.result)

		if tc.equal {
			testutil.CheckVal(t, wantStr, gotStr)
		} else {
			testutil.CheckNotEqual(t, wantStr, gotStr)
		}
	}
}

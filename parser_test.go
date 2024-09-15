package leetkit

import (
	"fmt"
	"testing"

	"github.com/alex-telpis/leetkit/internal/testutil"
)

func TestParse(t *testing.T) {
	{
		input := `"123"`
		v, err := tryParse[string](input)
		testutil.CheckNilErr(t, err)
		testutil.CheckVal(t, "123", v)
		testutil.CheckVal(t, input, Sprint(v))
	}
	{
		input := "123"
		v, err := tryParse[int](input)
		testutil.CheckNilErr(t, err)
		testutil.CheckVal(t, 123, v)
		testutil.CheckVal(t, input, Sprint(v))
	}
	{
		input := `"1"`
		v, err := tryParse[byte](input)
		testutil.CheckNilErr(t, err)
		testutil.CheckVal(t, byte('1'), v)
		testutil.CheckVal(t, input, Sprint(v))
	}
	{
		input := `1.2`
		v, err := tryParse[float64](input)
		testutil.CheckNilErr(t, err)
		testutil.CheckVal(t, float64(1.2), v)
	}
	{
		input := `true`
		v, err := tryParse[bool](input)
		testutil.CheckNilErr(t, err)
		testutil.CheckVal(t, true, v)
		testutil.CheckVal(t, input, Sprint(v))
	}
	{
		input := `false`
		v, err := tryParse[bool](input)
		testutil.CheckNilErr(t, err)
		testutil.CheckVal(t, false, v)
		testutil.CheckVal(t, input, Sprint(v))
	}
	{
		input := `[]`
		v, err := tryParse[[]int](input)
		testutil.CheckNilErr(t, err)
		testutil.CheckSlice(t, []int{}, v)
		testutil.CheckVal(t, input, Sprint(v))
	}
	{
		input := `[1,2,3]`
		v, err := tryParse[[]int](input)
		testutil.CheckNilErr(t, err)
		testutil.CheckSlice(t, []int{1, 2, 3}, v)
		testutil.CheckVal(t, input, Sprint(v))
	}
	{
		input := `["1","2","3"]`
		v, err := tryParse[[]string](input)
		testutil.CheckNilErr(t, err)
		testutil.CheckSlice(t, []string{"1", "2", "3"}, v)
		testutil.CheckVal(t, input, Sprint(v))
	}
	{
		input := `[[1,2,3],[4,5]]`
		v, err := tryParse[[][]int](input)
		testutil.CheckNilErr(t, err)
		testutil.CheckVal(t, 2, len(v))
		testutil.CheckSlice(t, []int{4, 5}, v[1])
		testutil.CheckVal(t, input, Sprint(v))
	}
	{
		input := `[1,2,3]`
		v, err := tryParse[*TreeNode](input)
		testutil.CheckNilErr(t, err)
		testutil.CheckVal(t, 1, v.Val)
		testutil.CheckVal(t, 2, v.Left.Val)
		testutil.CheckVal(t, 3, v.Right.Val)
		testutil.CheckVal(t, input, Sprint(v))
	}
	{
		input := `[1,2,3]`
		v, err := tryParse[*ListNode](input)
		testutil.CheckNilErr(t, err)
		testutil.CheckVal(t, 1, v.Val)
		testutil.CheckVal(t, 2, v.Next.Val)
		testutil.CheckVal(t, 3, v.Next.Next.Val)
		testutil.CheckVal(t, input, Sprint(v))
	}
}

func TestSplit(t *testing.T) {
	tests := []struct {
		input   string
		want    []string
		wantErr bool
	}{
		{"[]", []string{}, false},
		{"[1]", []string{"1"}, false},
		{"[[1], [2]]", []string{"[1]", "[2]"}, false},
		{"[1,2,3]", []string{"1", "2", "3"}, false},
		{"[1, 2, 3]", []string{"1", "2", "3"}, false},
		{" [1,2,3] ", []string{"1", "2", "3"}, false},
		{`[1, "2, 3"]`, []string{"1", `"2, 3"`}, false},
		{`[1,2,"[","[]"]`, []string{"1", "2", `"["`, `"[]"`}, false},
		{`[null,1,2,null]`, []string{"null", "1", "2", "null"}, false},
		{"[1,2,3,]", nil, true},
		{"[1,2,3", nil, true},
		{"1,2,3", nil, true},
	}

	for _, tc := range tests {
		got, err := split(tc.input)
		if tc.wantErr && err == nil {
			t.Error("expected error")
		}

		if !tc.wantErr {
			testutil.CheckNilErr(t, err)
		}

		testutil.CheckSlice(t, tc.want, got)
	}
}

func TestSprintNaryTree(t *testing.T) {
	type testcase struct {
		tree *NaryTreeNode
		want string
	}
	tests := []testcase{
		{
			tree: &NaryTreeNode{Val: 1},
			want: "[1]",
		},
		{
			tree: &NaryTreeNode{
				Val: 1,
				Children: []*NaryTreeNode{
					{
						Val: 3,
						Children: []*NaryTreeNode{
							{Val: 5},
							{Val: 6},
						},
					},
					{Val: 2},
					{Val: 4},
				},
			},
			want: "[1,null,3,2,4,null,5,6]",
		},
		{
			tree: &NaryTreeNode{
				Val: 1,
				Children: []*NaryTreeNode{
					{Val: 2},
					{
						Val: 3,
						Children: []*NaryTreeNode{
							{Val: 6},
							{
								Val: 7,
								Children: []*NaryTreeNode{
									{
										Val:      11,
										Children: []*NaryTreeNode{{Val: 14}},
									},
								},
							},
						},
					},
					{
						Val: 4,
						Children: []*NaryTreeNode{
							{Val: 8, Children: []*NaryTreeNode{{Val: 12}}},
						},
					},
					{
						Val: 5,
						Children: []*NaryTreeNode{
							{
								Val:      9,
								Children: []*NaryTreeNode{{Val: 13}},
							},
							{Val: 10},
						},
					},
				},
			},
			want: "[1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]",
		},
	}
	for _, test := range tests {
		t.Run(
			"", func(t *testing.T) {
				testutil.CheckVal(t, test.want, test.tree.String())
			},
		)
	}
}

func TestParseNaryTree(t *testing.T) {
	testcases := []string{
		"[1]",
		"[1,null,3,2,4,null,5,6]",
		"[1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]",
	}
	for _, test := range testcases {
		t.Run(
			"", func(t *testing.T) {
				tree, err := tryParseNaryTreeNode(test)
				testutil.CheckNilErr(t, err)
				testutil.CheckVal(t, test, tree.String())
			},
		)
	}
}

func TestCycleDetection(t *testing.T) {
	linkedList := &ListNode{Val: 1}
	linkedList.Next = &ListNode{Val: 2, Next: linkedList}

	tree := &TreeNode{Val: 1}
	tree.Left = &TreeNode{Val: 2, Right: tree}

	naryTree := &NaryTreeNode{Val: 1}
	naryTree.Children = []*NaryTreeNode{{Val: 2, Children: []*NaryTreeNode{naryTree}}}

	tests := []fmt.Stringer{
		linkedList,
		tree,
		naryTree,
	}

	for _, tc := range tests {
		testutil.CheckPanic(t, func() { tc.String() })
	}
}

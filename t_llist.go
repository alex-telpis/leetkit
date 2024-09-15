package leetkit

import (
	"encoding/json"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	seen := make(map[*ListNode]bool, 10)

	sb := &strings.Builder{}
	sb.WriteByte('[')
	for ; l != nil; l = l.Next {
		if sb.Len() > 1 {
			sb.WriteByte(',')
		}
		sb.WriteString(strconv.Itoa(l.Val))

		if seen[l] {
			panic(errCycle)
		}
		seen[l] = true
	}
	sb.WriteByte(']')
	return sb.String()
}

// ParseListNode is a shorthand for Parse[*ListNode](s).
func ParseListNode(s string) *ListNode { return panicOnParseErr(s, tryParseListNode) }

func tryParseListNode(s string) (*ListNode, error) {
	var res []*int
	err := json.Unmarshal([]byte(s), &res)
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, nil
	}
	root := &ListNode{}
	n := root
	for i := 0; i < len(res)-1; i++ {
		n.Val = *res[i]
		n.Next = &ListNode{}
		n = n.Next
	}
	n.Val = *res[len(res)-1]
	return root, nil
}

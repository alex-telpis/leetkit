package leetkit

import (
	"encoding/json"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) String() string {
	nodes := []*TreeNode{}
	queue := []*TreeNode{t}
	seen := make(map[*TreeNode]bool, 10)
	for len(queue) > 0 {
		t, queue = queue[0], queue[1:]
		nodes = append(nodes, t)
		if t != nil {
			if seen[t] {
				panic(errCycle)
			}
			seen[t] = true

			queue = append(queue, t.Left, t.Right)
		}
	}

	for len(nodes) > 0 && nodes[len(nodes)-1] == nil {
		nodes = nodes[:len(nodes)-1]
	}

	sb := &strings.Builder{}
	sb.WriteByte('[')
	for _, node := range nodes {
		if sb.Len() > 1 {
			sb.WriteByte(',')
		}
		if node != nil {
			sb.WriteString(strconv.Itoa(node.Val))
		} else {
			sb.WriteString("null")
		}
	}
	sb.WriteByte(']')

	return sb.String()
}

// ParseTreeNode is a shorthand for Parse[*TreeNode](s).
func ParseTreeNode(s string) *TreeNode { return panicOnParseErr(s, tryParseTreeNode) }

func tryParseTreeNode(s string) (*TreeNode, error) {
	var res []*int
	err := json.Unmarshal([]byte(s), &res)
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, nil
	}
	nodes := make([]*TreeNode, len(res))
	for i := 0; i < len(res); i++ {
		if res[i] != nil {
			nodes[i] = &TreeNode{Val: *res[i]}
		}
	}
	root := nodes[0]
	for i, j := 0, 1; j < len(res); i++ {
		if nodes[i] != nil {
			nodes[i].Left = nodes[j]
			j++
			if j >= len(res) {
				break
			}
			nodes[i].Right = nodes[j]
			j++
			if j >= len(res) {
				break
			}
		}
	}
	return root, nil
}

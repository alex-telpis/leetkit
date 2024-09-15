package leetkit

import (
	"encoding/json"
	"strconv"
	"strings"
)

type NaryTreeNode struct {
	Val      int
	Children []*NaryTreeNode
}

func (t *NaryTreeNode) String() string {
	nodes := []*NaryTreeNode{}
	q := []*NaryTreeNode{{Children: []*NaryTreeNode{t}}}
	seen := make(map[*NaryTreeNode]bool, 10)

	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		nodes = append(nodes, node)

		if node != nil {
			if seen[node] {
				panic(errCycle)
			}
			seen[node] = true

			if len(node.Children) > 0 {
				q = append(q, node.Children...)
			}
			q = append(q, nil)
		}
	}

	nodes = nodes[1:]

	for len(nodes) > 0 && nodes[len(nodes)-1] == nil {
		nodes = nodes[:len(nodes)-1]
	}

	sb := strings.Builder{}
	sb.WriteByte('[')
	for _, node := range nodes {
		if sb.Len() > 1 {
			sb.WriteByte(',')
		}
		if node == nil {
			sb.WriteString("null")
		} else {
			sb.WriteString(strconv.Itoa(node.Val))
		}
	}
	sb.WriteByte(']')
	return sb.String()
}

// ParseNaryTreeNode is a shorthand for Parse[*NaryTreeNode](s).
func ParseNaryTreeNode(s string) *NaryTreeNode { return panicOnParseErr(s, tryParseNaryTreeNode) }

func tryParseNaryTreeNode(s string) (*NaryTreeNode, error) {
	var res []*int
	if err := json.Unmarshal([]byte(s), &res); err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, nil
	}

	root := &NaryTreeNode{}
	q := []*NaryTreeNode{root}
	for i := 0; i < len(res); i++ {
		node := q[0]
		q = q[1:]
		for ; i < len(res) && res[i] != nil; i++ {
			n := &NaryTreeNode{Val: *res[i]}
			node.Children = append(node.Children, n)
			q = append(q, n)
		}
	}

	return root.Children[0], nil
}

# LeetKit for Go [![Go Reference](https://pkg.go.dev/badge/github.com/alex-telpis/leetkit.svg)](https://pkg.go.dev/github.com/alex-telpis/leetkit) [![Go Report Card](https://goreportcard.com/badge/github.com/alex-telpis/leetkit)](https://goreportcard.com/report/github.com/alex-telpis/leetkit)

So, you've decided to take the **Go**-ing-against-the-grain route on LeetCode?
Well, LeetKit is here to make your life a little less painful.
This library simplifies working with complex data structures like
linked lists, trees, matrices, providing common LeetCode type definitions
and functions for parsing example inputs.

## Installation

```bash
go get -u github.com/alex-telpis/leetkit
```

## Basic Usage

0. ~~Consider using Python.~~
1. Define a struct in your solution using a type alias, e.g., `type TreeNode = leetkit.TreeNode`.
2. Use the `Parse*` functions to handle Leetcode input samples.
3. Use `Verify()` to compare your result with the expected value and print the outcome.

Example:

```go
// https://leetcode.com/problems/linked-list-in-binary-tree
package main

import "github.com/alex-telpis/leetkit"


// Define local type aliases for TreeNode and ListNode to match LeetCode's signature.
type TreeNode = leetkit.TreeNode
type ListNode = leetkit.ListNode

func main() {
    ll := leetkit.ParseListNode("[1,4,2,6]")
    bt := leetkit.ParseTreeNode("[1,4,4,null,2,2,null,1,null,6,8,null,null,null,null,1,3]")

    leetkit.Verify(true, isSubPath(ll, bt))
}

// Your submission code begins here:

func isSubPath(head *ListNode, root *TreeNode) bool {
    // ... solution

    return true
}
```

## Matrices

`Parse<Int | String | Byte>Matrix` functions are available as convenient shortcuts
to the generic `Parse()` function. The following example also demonstrates
how the `Verify()` function works with the string representation of the expected
result.

Example:

```go
// https://leetcode.com/problems/spiral-matrix
package main

import (
    "slices"

    "github.com/alex-telpis/leetkit"
)

func main() {
    leetkit.Verify(
        "[1,2,3,6,9,8,7,4,5]",
        spiralOrder(leetkit.ParseIntMatrix("[[1,2,3],[4,5,6],[7,8,9]]")),
    )
    leetkit.Verify(
        "[1,2,3,4,8,12,11,10,9,5,6,7]",
        spiralOrder(leetkit.ParseIntMatrix("[[1,2,3,4],[5,6,7,8],[9,10,11,12]]")),
    )
    leetkit.Verify(
        "[1,2,3,4,8,12,16,15,14,13,9,5,6,7,11,10]",
        spiralOrder(leetkit.ParseIntMatrix("[[1,2,3,4],[5,6,7,8],[9,10,11,12],[13,14,15,16]]")),
    )
}

// Your submission code begins here:

func spiralOrder(matrix [][]int) []int {
    rMin, rMax := 0, len(matrix)-1
    cMin, cMax := 0, len(matrix[0])-1

    getCol := func(matrix [][]int, i int) []int {
        res := make([]int, 0, len(matrix))
        for r := rMin; r <= rMax; r++ {
            res = append(res, matrix[r][i])
        }

        return res
    }
    res := make([]int, 0, len(matrix)*len(matrix[0]))

    for i := 0; len(res) < cap(res); i++ {
        switch i % 4 {
        case 0: // left to right
            res = append(res, matrix[rMin][cMin:cMax+1]...)
            rMin++

        case 1: // top to bottom
            res = append(res, getCol(matrix, cMax)...)
            cMax--

        case 2: // right to leftrow
            row := matrix[rMax][cMin : cMax+1]
            slices.Reverse(row)
            res = append(res, row...)
            rMax--

        case 3: // bottom to up
            row := getCol(matrix, 0)
            slices.Reverse(row)
            res = append(res, row...)
            cMin++
        }
    }

    return res
}
```

## Heaps

Trying to use a heap in Go is a significant time investment 🥲.
To quickly validate your heap-based solution, you can use the built-in `IntMinHeap`
and `IntMaxHeap`. Just don't forget to include the heap in your LeetCode submission.

Example:

```go
// https://leetcode.com/problems/last-stone-weight
package main

import (
    "container/heap"

    "github.com/alex-telpis/leetkit"
)

func main() {
    leetkit.Verify(1, lastStoneWeight([]int{2, 7, 4, 1, 8, 1}))
}


// Your submission code begins here:

func lastStoneWeight(stones []int) int {
    h := &leetkit.IntMaxHeap{} // use temporary implementation
    *h= stones
    heap.Init(h)

    for {
        switch h.Len() {
        case 0:
            return 0

        case 1:
            return (*h)[0]

        default:
            s1 := heap.Pop(h).(int)
            s2 := heap.Pop(h).(int)

            if s1 > s2 {
                heap.Push(h, s1-s2)
            }
        }
    }
}
```

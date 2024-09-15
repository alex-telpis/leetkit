# LeetKit for Go

So, you've decided to take the **Go**-ing-against-the-grain route on LeetCode?
Well, LeetKit is here to make your life a little less painful.
This library simplifies working with complex data structures like
linked lists, trees, matrices, providing common LeetCode type definitions
and functions for parsing example inputs.

## Usage

0. ~~Consider using Python.~~
1. Define a struct in your solution using a type alias, e.g., `type TreeNode = leetkit.TreeNode`.
2. Use the `Parse*` functions to handle Leetcode input samples.
3. Use `Verify()` to compare your result with the expected value and print the outcome.

Example:

```go
// https://leetcode.com/problems/linked-list-in-binary-tree/description/
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
    return false
}
```

## Installation

```bash
go get -u github.com/alex-telpis/leetkit
```

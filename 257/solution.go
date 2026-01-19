package main

import (
	"fmt"
	"strconv"
)

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	visited := map[*TreeNode]bool{}
	retval := []string{}
	dfsRecursive(root, visited, "", &retval)
	return retval
}

func dfsRecursive(node *TreeNode, visited map[*TreeNode]bool, curStr string, strs *[]string) {
	fmt.Println("currentNode: ", node)
	if visited[node] == false {
		visited[node] = true

		newStr := curStr
		if len(curStr) > 0 {
			newStr += "->"
		}

		newStr += strconv.Itoa(node.Val)
		if node.Left == nil && node.Right == nil {
			fmt.Println("appending", newStr)
			*strs = append(*strs, newStr)
			return
		}

		if node.Left != nil {
			dfsRecursive(node.Left, visited, newStr, strs)
		}
		if node.Right != nil {
			dfsRecursive(node.Right, visited, newStr, strs)
		}
	}
}

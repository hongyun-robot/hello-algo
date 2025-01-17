// File: preorder_traversal_i_compact.go
// Created Time: 2023-05-09
// Author: Reanon (793584285@qq.com)

package chapter_backtracking

import (
	. "github.com/krahets/hello-algo/pkg"
)

/* 前序遍历：例题三 */
func preOrderIII(root *TreeNode, res *[][]*TreeNode, path *[]*TreeNode) {
	// 剪枝
	if root == nil || root.Val == 3 {
		return
	}
	// 尝试
	*path = append(*path, root)
	if int(root.Val) == 7 {
		// 记录解
		*res = append(*res, *path)
		return
	}
	preOrderIII(root.Left, res, path)
	preOrderIII(root.Right, res, path)
	// 回退
	*path = (*path)[:len(*path)-1]
}

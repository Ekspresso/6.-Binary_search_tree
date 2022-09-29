// Функция переводит отсортированный массив в бинарное дерево поиска.

package binarysearchtree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) > 0 {
		n := len(nums) / 2
		root := new(TreeNode)
		root.Val = nums[n]
		root.Left = sortedArrayToBST(nums[:n])
		root.Right = sortedArrayToBST(nums[n+1:])
		return root
	}
	return nil
}

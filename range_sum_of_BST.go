// Функция выводит сумму элементов бинарного дерева поиска
// из заданного диапазона.

package binarysearchtree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	sum := 0
	if root != nil {
		if root.Val > low && root.Val <= high {
			sum += rangeSumBST(root.Left, low, high)
			sum += rangeSumBST(root.Right, low, high)
		} else if root.Val > high {
			sum += rangeSumBST(root.Left, low, high)
		} else if root.Val <= low {
			sum += rangeSumBST(root.Right, low, high)
		}
		if root.Val >= low && root.Val <= high {
			return sum + root.Val
		}
	}
	return sum
}

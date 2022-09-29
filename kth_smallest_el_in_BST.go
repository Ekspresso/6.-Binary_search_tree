// Функция возвращает k-тое по величине значение в бинарном дереве поиска.
// The number of nodes in the tree is n.
// 1 <= k <= n <= 10^4
// 0 <= Node.val <= 10^4

package binarysearchtree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Переменная isZero нужна для корректного вывода,
// если искомый элемент равен 0.
func kthSmallest(root *TreeNode, k int) int {
	isZero := false
	ret, isZero := outNegativeSearch(root, k, isZero)
	return -ret
}

func outNegativeSearch(root *TreeNode, k int, isZero bool) (int, bool) {
	if root != nil {
		k, isZero = outNegativeSearch(root.Left, k, isZero)
		// Условие для выхода из рекурсий с найденым значением.
		if k < 0 || isZero {
			return k, isZero
		}
		k -= 1
		// Если k == 0 или флаг поиска нуля false,
		// значит текущее значение искомое, и оно
		// записывается в k отрицательным значением
		// или, если это ноль, ставится флаг isZero в положение true.
		if k == 0 && !isZero {
			if root.Val == 0 {
				isZero = true
			}
			return -root.Val, isZero
		}
		k, isZero = outNegativeSearch(root.Right, k, isZero)
	}
	return k, isZero
}

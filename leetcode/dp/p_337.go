package dp

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var money map[*TreeNode]int

func rob3(root *TreeNode) int {
	money = map[*TreeNode]int{}
	return tryRob(root)
}

func tryRob(root *TreeNode) int {
	if val, ok := money[root]; ok {
		return val
	}
	a := root.Val
	b := 0
	if root.Left != nil {
		b += tryRob(root.Left)
		if root.Left.Left != nil {
			a += tryRob(root.Left.Left)
		}
		if root.Left.Right != nil {
			a += tryRob(root.Left.Right)
		}
	}
	if root.Right != nil {
		b += tryRob(root.Right)
		if root.Right.Left != nil {
			a += tryRob(root.Right.Left)
		}
		if root.Right.Right != nil {
			a += tryRob(root.Right.Right)
		}
	}
	money[root] = max(a, b)
	return money[root]
}

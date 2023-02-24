package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil {
		if q == nil {
			return true
		}
		return false
	} else if q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	same := true
	if p.Left != nil {
		if q.Left == nil {
			return false
		}
		same = same && isSameTree(p.Left, q.Left)
	} else if q.Left != nil {
		return false
	}
	if p.Right != nil {
		if q.Right == nil {
			return false
		}
		same = same && isSameTree(p.Right, q.Right)
	} else if q.Right != nil {
		return false
	}
	return same
}

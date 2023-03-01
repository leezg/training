package tree

//TODO: rbTree

type Color bool

const (
	RED   Color = true
	BLACK Color = false
)

type RedBlackTree struct {
	Root *Node
}

func (t *RedBlackTree) Insert(key int, value interface{}) {
	t.Root = insert(t.Root, key, value)
	t.Root.Color = BLACK
}

func (t *RedBlackTree) Search(key int) (interface{}, bool) {
	n := t.Root
	for n != nil {
		if key < n.Key {
			n = n.Left
		} else if key > n.Key {
			n = n.Right
		} else {
			return n.Value, true
		}
	}
	return nil, false
}

type Node struct {
	Key         int
	Value       interface{}
	Color       Color
	Left, Right *Node
}

func (n *Node) IsRed() bool {
	if n == nil {
		return false
	}
	return n.Color == RED
}

func insert(n *Node, key int, value interface{}) *Node {
	if n == nil {
		return &Node{Key: key, Value: value, Color: RED}
	}
	if key < n.Key {
		n.Left = insert(n.Left, key, value)
	} else if key > n.Key {
		n.Right = insert(n.Right, key, value)
	} else {
		n.Value = value
	}

	if n.Right.IsRed() && !n.Left.IsRed() {
		n = rotateLeft(n)
	}
	if n.Left.IsRed() && n.Left.Left.IsRed() {
		n = rotateRight(n)
	}
	if n.Left.IsRed() && n.Right.IsRed() {
		flipColors(n)
	}

	return n
}

func rotateLeft(h *Node) *Node {
	x := h.Right
	h.Right = x.Left
	x.Left = h
	x.Color = h.Color
	h.Color = RED
	return x
}

func rotateRight(h *Node) *Node {
	x := h.Left
	h.Left = x.Right
	x.Right = h
	x.Color = h.Color
	h.Color = RED
	return x
}

func flipColors(h *Node) {
	h.Color = RED
	h.Left.Color = BLACK
	h.Right.Color = BLACK
}

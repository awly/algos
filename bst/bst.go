// Binary search tree
package bst

// type for elements of the tree
type Elem interface {
	Cmp(Elem) int
}

// wrapper for ints that satisfies Elem interface
type Int int

func (i Int) Cmp(other Elem) int {
	j := other.(Int)
	return int(i) - int(j)
}

// element in the tree
type Node struct {
	parent, left, right *Node

	Val Elem
}

// create new tree. same as simply using &Node{}
func NewBST() *Node {
	return &Node{}
}

// insert new element into the tree
// returns false if such element already exists (Cmp returns 0)
func (n *Node) Insert(k Elem) bool {
	// if this is first insert
	if n.Val == nil {
		n.Val = k
		return true
	}

	cmp := k.Cmp(n.Val)
	if cmp < 0 {
		if n.left == nil {
			n.left = &Node{Val: k, parent: n}
			return true
		} else {
			return n.left.Insert(k)
		}
	} else if cmp > 0 {
		if n.right == nil {
			n.right = &Node{Val: k, parent: n}
			return true
		} else {
			return n.right.Insert(k)
		}
	}
	return false
}

// returns the node with matching value
func (n *Node) Find(k Elem) *Node {
	cmp := k.Cmp(n.Val)
	if cmp == 0 {
		return n
	} else if cmp < 0 {
		if n.left == nil {
			return nil
		}
		return n.left.Find(k)
	} else {
		if n.right == nil {
			return nil
		}
		return n.right.Find(k)
	}
}

// removes element from the tree
// returns false if element is not in the tree
func (n *Node) Remove(k Elem) bool {
	rm := n.Find(k)
	if rm == nil {
		return false
	}

	if rm.right != nil {
		next := rm.right
		for ; next.left != nil; next = next.left {
		}
		rm.Val, next.Val = next.Val, rm.Val
		next.parent.left = nil
	} else if rm.left != nil {
		prev := rm.left
		for ; prev.right != nil; prev = prev.right {
		}
		rm.Val, prev.Val = prev.Val, rm.Val
		prev.parent.right = nil
	} else {
		n.Val = nil
	}

	return true
}

// calls the tree and calls fn for each element in-order
func (n *Node) Walk(fn func(Elem)) {
	if n.left != nil {
		n.left.Walk(fn)
	}
	fn(n.Val)
	if n.right != nil {
		n.right.Walk(fn)
	}
}

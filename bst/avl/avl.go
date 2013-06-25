// AVL tree (ballanced BST)
// measures time to build the tree for 1Mil elements
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	data := make([]int, 1e6)
	for i := 0; i < len(data); i++ {
		data[i] = rand.Int()
	}
	start := time.Now()
	tree := NewAVLTree(data)
	fmt.Println(time.Since(start))
	tree.traverse(func(n *node) {
		// uncomment to print sorted tree
		//fmt.Print(n.Val, " ")
	})
}

type node struct {
	parent, Left, Right *node
	Val, size           int
}

func NewAVLTree(data []int) *node {
	root := &node{Val: data[0]}
	for i := 1; i < len(data); i++ {
		root = root.insert(data[i])
	}
	return root
}

func (n *node) insert(k int) *node {
	if k < n.Val {
		if n.Left == nil {
			n.Left = &node{Val: k, size: 0, parent: n}
		} else {
			n.Left.insert(k)
		}
	} else {
		if n.Right == nil {
			n.Right = &node{Val: k, size: 0, parent: n}
		} else {
			n.Right.insert(k)
		}
	}
	n.size++
	n.fixAVL()
	if n.parent != nil {
		return n.parent
	}
	return n
}

func (n *node) find(k int) *node {
	if k == n.Val {
		return n
	} else if k < n.Val {
		if n.Left == nil {
			return nil
		}
		return n.Left.find(k)
	} else {
		if n.Right == nil {
			return nil
		}
		return n.Right.find(k)
	}
}

func (n *node) rotLeft() {
	if n.parent != nil {
		if n.parent.Left == n {
			n.parent.Left = n.Right
		} else {
			n.parent.Right = n.Right
		}
	}
	n.parent, n.Right.parent = n.Right, n.parent
	n.Right, n.Right.Left = n.Right.Left, n
	if n.Right != nil {
		n.Right.parent = n
	}
}

func (n *node) rotRight() {
	if n.parent != nil {
		if n.parent.Left == n {
			n.parent.Left = n.Left
		} else {
			n.parent.Right = n.Left
		}
	}
	n.parent, n.Left.parent = n.Left, n.parent
	n.Left, n.Left.Right = n.Left.Right, n
	if n.Left != nil {
		n.Left.parent = n
	}
}

func (n *node) balanceFactor() int {
	rsize, lsize := -1, -1
	if n.Left != nil {
		lsize = n.Left.size
	}
	if n.Right != nil {
		rsize = n.Right.size
	}
	return lsize - rsize
}

func (n *node) fixAVL() {
	bf := n.balanceFactor()
	if bf > 1 {
		if n.Left.balanceFactor() < 0 {
			n.Left.rotLeft()
		}
		n.rotRight()
	} else if bf < -1 {
		if n.Right.balanceFactor() > 0 {
			n.Right.rotRight()
		}
		n.rotLeft()
	}
}

func (n *node) traverse(fn func(*node)) {
	if n.Left != nil {
		n.Left.traverse(fn)
	}
	fn(n)
	if n.Right != nil {
		n.Right.traverse(fn)
	}
}

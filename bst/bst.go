// Binary search tree
// measures time to build the tree of 1Mil elements
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	data := make([]int, 1e6)
	for i, _ := range data {
		data[i] = rand.Int()
	}
	start := time.Now()
	tree := NewBST(data)
	fmt.Println(time.Since(start))
	tree.traverse(func(n *node) {
		// uncomment to print sorted slice
		//fmt.Print(n.val, " ")
	})
}

type node struct {
	parent, left, right *node
	val, size           int
}

func NewBST(data []int) *node {
	root := &node{val: data[0]}
	for i := 1; i < len(data); i++ {
		root.insert(data[i])
	}
	return root
}

func (n *node) insert(k int) bool {
	if k < n.val {
		if n.left == nil {
			n.left = &node{val: k, size: 1, parent: n}
			n.size++
			return true
		} else {
			res := n.left.insert(k)
			if res {
				n.size++
			}
			return res
		}
	} else if k > n.val {
		if n.right == nil {
			n.right = &node{val: k, size: 1, parent: n}
			n.size++
			return true
		} else {
			res := n.right.insert(k)
			if res {
				n.size++
			}
			return res
		}
	}
	return false
}

func (n *node) find(k int) *node {
	if k == n.val {
		return n
	} else if k < n.val {
		if n.left == nil {
			return nil
		}
		return n.left.find(k)
	} else {
		if n.right == nil {
			return nil
		}
		return n.right.find(k)
	}
}

func (n *node) traverse(fn func(*node)) {
	if n.left != nil {
		n.left.traverse(fn)
	}
	fn(n)
	if n.right != nil {
		n.right.traverse(fn)
	}
}

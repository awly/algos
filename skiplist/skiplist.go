// skiplist: implementation of skip lists
// list elements store key/value pairs
// key has to satisfy Comparer interface
package skiplist

import (
	"fmt"
	"math/rand"
)

// interface for list keys
type Comparer interface {
	Cmp(Comparer) int
}

// example Comparer implementation for ints
type Int int

func (i Int) Cmp(j Comparer) int {
	if jint, ok := j.(Int); ok {
		return int(i) - int(jint)
	} else {
		panic("Cmp: can't compare Int with " + fmt.Sprintf("%T", j))
	}
}

// Slist implements skip list. zero-value is usable
type Slist struct {
	top, bottom *node // top - top-most list, bottom - lowest list, containing all elements
}

type node struct {
	// next and previous elements in the same list (same level)
	// upper and lower elements from lists above and below
	next, prev, down, up *node
	// key and value of element
	key Comparer
	val interface{}
}

// Insert adds new key/value pair to the list, preserving sorted order
// key can not be nil
// returns false if element is already in the list (if key.Cmp for at least one existing element returned 0), true otherwise
func (sl *Slist) Insert(k Comparer, v interface{}) bool {
	if sl.bottom == nil { // first element
		sl.bottom = &node{key: k, val: v}
		sl.top = sl.bottom
		return true
	}
	for n := sl.top; n != nil; {
		cmp := k.Cmp(n.key)
		switch {
		case cmp == 0:
			return false // element already in list
		case cmp > 0:
			if n.next != nil { // go forward
				n = n.next
			} else if n.down != nil { // go lower
				n = n.down
			} else { // no elements in front or below, insert
				newn := &node{key: k, val: v, prev: n}
				n.next = newn
				sl.promote(newn)
				return true
			}
		case cmp < 0:
			if n.prev != nil {
				if n.prev.down != nil {
					n = n.prev.down
				} else {
					newn := &node{next: n, prev: n.prev, key: k, val: v}
					n.prev.next = newn
					n.prev = newn
					sl.promote(newn)
					return true
				}
			} else {
				// first element of the list is always the smallest
				// since we insert something smaller than 1st element, add new 1st element to every list
				newn := &node{next: sl.bottom, key: k, val: v}
				sl.bottom.prev = newn
				sl.bottom = newn
				if sl.bottom.next.up != nil {
					for newn = newn.next.up; newn != nil; newn = newn.next.up {
						newn = &node{next: newn, key: k, val: v, down: newn.down}
						newn.next.prev = newn
					}
				}
				sl.top = newn
				return true
			}
		}
	}
	return true
}

// Find returns value of the corresponding key in the list or nil if key is not present
func (sl *Slist) Find(k Comparer) interface{} {
	if sl.bottom == nil {
		return nil
	}
	for n := sl.top; n != nil; {
		cmp := k.Cmp(n.key)
		switch {
		case cmp == 0:
			return n.val
		case cmp > 0:
			if n.next != nil { // go forward
				n = n.next
			} else if n.down != nil { // go lower
				n = n.down
			} else {
				return nil
			}
		case cmp < 0:
			if n.prev != nil {
				if n.prev.down != nil {
					n = n.prev.down
				} else {
					return nil
				}
			} else {
				return nil
			}
		}
	}
	return nil
}

// promote randomly promotes node argument to the list above, creating it if necessary
// probability 50%
func (sl *Slist) promote(n *node) {
	var prevn, newn *node
	for rand.Float64() > 0.5 {
		newn = &node{down: n, key: n.key, val: n.val}
		n.up = newn
		for prevn = n.prev; prevn.prev != nil && prevn.up == nil; prevn = prevn.prev {
		}
		if prevn.up == nil { // we're at top list, create new tier
			prevn.up = &node{down: prevn, next: newn, key: prevn.key, val: prevn.val}
			newn.prev = prevn.up
			sl.top = prevn.up
		} else {
			prevn.up.next, newn.next, newn.prev = newn, prevn.up.next, prevn.up
			if newn.next != nil {
				newn.next.prev = newn
			}
		}
		n = newn
	}
}

// Walk calls fn for evey element of list in sorted order and stops if fn returns false
func (sl *Slist) Walk(fn func(Comparer, interface{}) bool) {
	for n := sl.bottom; n != nil; n = n.next {
		if !fn(n.key, n.val) {
			return
		}
	}
}

// print prints all lists from top-most (the shortest one) til lowest
func (sl *Slist) print() {
	for st := sl.top; st != nil; st = st.down {
		for n := st; n != nil; n = n.next {
			fmt.Print(n.key, " ")
		}
		fmt.Println()
	}
	fmt.Println("--------------------------")
}

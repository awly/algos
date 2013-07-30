package bst

import (
	"testing"
)

var data = []Int{5, 3, 7, 2, 8, 1, 6, 4}

func TestInsertRoot(t *testing.T) {
	tree := NewBST()
	if !tree.Insert(Int(1)) {
		t.Fatal("failed to insert root value")
	}
}

func TestInsert(t *testing.T) {
	tree := NewBST()
	for _, e := range data {
		if !tree.Insert(e) {
			t.Fatal("failed to insert value", e)
		}
	}
}

func TestFind(t *testing.T) {
	tree := NewBST()
	for _, e := range data {
		if !tree.Insert(e) {
			t.Fatal("failed to insert value", e)
		}
	}
	if tree.Find(data[3]) == nil {
		t.Fatal("failed to find", data[3])
	}
}

func TestRemove(t *testing.T) {
	tree := NewBST()
	for _, e := range data {
		if !tree.Insert(e) {
			t.Fatal("failed to insert value", e)
		}
	}
	if tree.Find(data[3]) == nil {
		t.Fatal("failed to find", data[3])
	}
	if !tree.Remove(data[3]) {
		t.Fatal("failed to remove", data[3])
	}
	if tree.Find(data[3]) != nil {
		t.Fatal("element found in tree after removal", data[3])
	}
}

func TestSorted(t *testing.T) {
	tree := NewBST()
	for _, e := range data {
		if !tree.Insert(e) {
			t.Fatal("failed to insert value", e)
		}
	}
	var last Elem
	tree.Walk(func(e Elem) {
		if last != nil {
			if last.Cmp(e) > 0 {
				t.Fatal("tree is not sorted: ", e, " after ", last)
			}
		}
		last = e
		return
	})
}

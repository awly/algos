package skiplist

import (
	"math/rand"
	"testing"
)

func init() {
	rand.Seed(0)
}

var keys = []Int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
var vals = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}

func TestInsertIncr(t *testing.T) {
	sl := &Slist{}
	for i, v := range keys {
		if !sl.Insert(v, vals[i]) {
			t.Fatal("failed to add new element", v)
		}
	}
	checkkv(sl, t)
}

func TestInsertDecr(t *testing.T) {
	sl := &Slist{}
	for i := len(keys) - 1; i >= 0; i-- {
		if !sl.Insert(keys[i], vals[i]) {
			t.Fatal("failed to add new element", keys[i])
		}
	}
	checkkv(sl, t)
}

func TestInsertRand(t *testing.T) {
	sl := &Slist{}
	for _, i := range rand.Perm(len(keys)) {
		if !sl.Insert(keys[i], vals[i]) {
			t.Fatal("failed to add new element", keys[i])
		}
	}
	checkkv(sl, t)
}

func TestFindFromZero(t *testing.T) {
	sl := &Slist{}
	if v := sl.Find(Int(0)); v != nil {
		t.Fatal("find from empty list returned non-nil value", v)
	}
}

func TestFind(t *testing.T) {
	sl := &Slist{}
	for i, v := range keys {
		if !sl.Insert(v, vals[i]) {
			t.Fatal("failed to add new element", v)
		}
	}
	for i, v := range keys {
		val := sl.Find(v)
		if val == nil || val != vals[i] {
			t.Fatal("failed to find element for key", v, ": expected", keys[i], "got", val)
		}
	}
}

// check keys order, key-value correspondance and completeness of list
// presumes list being filled with keys/values slices
func checkkv(sl *Slist, t *testing.T) {
	i := 0
	sl.Walk(func(k Comparer, v interface{}) bool {
		if k != keys[i] {
			t.Fatal("order not preserved, expect", keys[i], "got", k)
			return false
		} else if v != vals[i] {
			t.Fatal("key-value correspondance not preserved, expect", keys[i], ":", vals[i], "got", k, ":", v)
			return false
		}
		i++
		return true
	})
	if i != len(keys) {
		t.Fatal("missing elements in the list, only", i, "found")
	}
}

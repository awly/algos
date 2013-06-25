// Hashtable
// measures time for:
// 1Mil random inserts
// 1Mil random retreives
// 1Mil random deletes
//
// for simplicity all random numbers are [1, 1e6)
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	ht := NewHT()
	start := time.Now()
	for i := 0; i < 1e6; i++ {
		k := rand.Intn(1e6)
		ht.put(k, i)
	}
	fmt.Println("inserts:", time.Since(start))
	start = time.Now()
	for i := 0; i < 1e6; i++ {
		k := rand.Intn(1e6)
		ht.get(k)
	}
	fmt.Println("retreives:", time.Since(start))
	start = time.Now()
	for i := 0; i < 1e6; i++ {
		k := rand.Intn(1e6)
		ht.del(k)
	}
	fmt.Println("deletes:", time.Since(start))
}

type buck struct {
	key  int
	val  int
	next *buck
}

type hashtable struct {
	data        []*buck
	count, size int
}

func NewHT() *hashtable {
	return &hashtable{data: make([]*buck, 8), size: 8}
}

func (ht *hashtable) put(key, val int) {
	hash := ht.hash(key)
	n := ht.data[hash]
	if n == nil {
		ht.data[hash] = &buck{key: key, val: val}
		ht.count++
	} else {
		for ; n.next != nil && n.key != key; n = n.next {
		}
		if n.key != key {
			n.next = &buck{key: key, val: val}
			ht.count++
		} else {
			n.val = val
		}
	}
	if ht.count >= ht.size {
		ht.grow()
	}
}

func (ht *hashtable) get(key int) (int, bool) {
	for n := ht.data[ht.hash(key)]; n != nil; n = n.next {
		if n.key == key {
			return n.val, true
		}
	}
	return 0, false
}

func (ht *hashtable) del(key int) (res bool) {
	hash := ht.hash(key)
	n := ht.data[hash]
	if n == nil {
		goto end
	}
	if n.key == key {
		ht.data[hash] = n.next
		ht.count--
		res = true
		goto end
	}
	for ; n.next != nil && n.next.key != key; n = n.next {
	}
	if n.next == nil {
		goto end
	}
	n.next = n.next.next
	ht.count--
	res = true
end:
	if ht.count < ht.size/4 && ht.size/4 >= 8 {
		ht.shrink()
	}
	return
}

func (ht *hashtable) grow() {
	ht.size *= 2
	ht.count = 0
	old := ht.data
	ht.data = make([]*buck, ht.size)
	for _, n := range old {
		for ; n != nil; n = n.next {
			ht.put(n.key, n.val)
		}
	}
}

func (ht *hashtable) shrink() {
	ht.size /= 4
	ht.count = 0
	old := ht.data
	ht.data = make([]*buck, ht.size)
	for _, n := range old {
		for ; n != nil; n = n.next {
			ht.put(n.key, n.val)
		}
	}
}

func (ht *hashtable) hash(i int) int {
	return i % ht.size
}

func (ht *hashtable) print() {
	for _, n := range ht.data {
		for ; n != nil; n = n.next {
			fmt.Print("(", n.key, ",", n.val, ") ")
		}
	}
	fmt.Println()
}

// orset: implementation of ordered set - a queue with unique elements
// the idea is to have a queue of elements and prevent dulpicates
//
// contains 2 implementations:
// - naive slow (O(n) insertion, no memory overhead) Slow version
// - cool fast (O(1) insertion, O(n) additional memory) hash-table + linked list version
package orset

// Keyer represents interface used for elements of ordered set
type Keyer interface {
	Key() int
}

type Int int

func (i Int) Key() int {
	return int(i)
}

// Ordered set, has to satisfy certain properties:
// - FIFO
// - no duplicates
// - adding duplicate element should move it to the end of queue
type Orset interface {
	PopFront() Keyer
	PushBack(Keyer)
}

type Slow []Keyer

func NewSlow() *Slow {
	s := make(Slow, 0)
	return &s
}

func (s *Slow) PopFront() Keyer {
	if len(*s) == 0 {
		return nil
	}
	front := (*s)[0]
	*s = (*s)[1:]
	return front
}

func (s *Slow) PushBack(k Keyer) {
	key := k.Key()
	for i, v := range *s {
		if v.Key() == key {
			*s = append((*s)[:i], (*s)[i+1:]...)
			break
		}
	}
	*s = append(*s, k)
}

type node struct {
	next, prev *node
	val        Keyer
}

type Fast struct {
	front, back *node
	keys        map[int]*node
}

func NewFast() *Fast {
	return &Fast{keys: make(map[int]*node)}
}

func (f *Fast) PopFront() Keyer {
	if f.front == nil {
		return nil
	}

	ret := f.front.val
	delete(f.keys, ret.Key())

	if f.front.next != nil {
		f.front.next.prev = nil
		f.front = f.front.next
	} else {
		f.front, f.back = nil, nil
	}

	return ret
}

func (f *Fast) PushBack(e Keyer) {
	key := e.Key()

	// remove existing duplicate if present
	if n, ok := f.keys[key]; ok {
		if n.next != nil && n.prev != nil {
			n.next.prev, n.prev.next = n.prev, n.next
		} else if n.next != nil && n.prev == nil {
			f.front = f.front.next
			f.front.prev = nil
		} else if n.next == nil && n.prev != nil {
			f.back = f.back.prev
			f.back.next = nil
		} else {
			f.front, f.back = nil, nil
		}
	}

	// append at the end
	if f.back != nil {
		f.back.next = &node{prev: f.back, val: e}
		f.back = f.back.next
	} else {
		f.front = &node{prev: f.back, val: e}
		f.back = f.front
	}
	f.keys[key] = f.back
}

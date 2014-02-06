package pmerge

import (
	"fmt"
	"runtime"
)

// Sort sorts data in-place using concurrent mergesort with
// GOMAXPROCS worker goroutines
func Sort(data []int) {
	if len(data) == 0 || len(data) == 1 {
		return
	}

	out := make(chan pair)
	in := make(chan block)

	blocks := make(map[int]block, len(data))
	for i := range data {
		blocks[i] = block{start: i, end: i + 1, data: data[i : i+1]}
	}

	for i := 0; i < runtime.GOMAXPROCS(-1); i++ {
		go merger(out, in)
	}

	for {
		p, ok := findPair(blocks)
		if ok {
			select {
			case out <- p:
				delete(blocks, p.a.start)
				delete(blocks, p.b.start)
				continue
			case nb := <-in:
				// if we have a pair pending, we should not be able to receive
				// complete sorted block (from 0 to len(data))
				blocks[nb.start] = nb
				continue
			}
			// try sending pair and receiving new block
		} else {
			nb := <-in
			if nb.start == 0 && nb.end == len(data) {
				// we're done, terminate workers
				close(out)
				return
			}
			blocks[nb.start] = nb
		}
	}
}

// merger listens on in chan for block pairs, merges them and sends
// merged block on out chan
func merger(in chan pair, out chan block) {
	for v := range in {
		//out <- merge(v.a, v.b)
		out <- mergeBuf(v.a, v.b)
	}
}

// merge takes two neighbour blocks and merges them in-place
// requires that blocks have the same underlying array
func merge(a, b block) block {
	if a.start == b.end {
		a, b = b, a
	}
	d := a.data[0 : b.end-a.start]
	p := b.start - a.start
	if p < 0 || p >= len(d) {
		panic(fmt.Sprintf("invalid arguments to merge: %v, %v", a, b))
	}
	i := 0
	for i < p && p < len(d) {
		if d[p] < d[i] {
			t := d[p]
			copy(d[i+1:p+1], d[i:p])
			d[i] = t
			p++
		}
		i++
	}

	return block{data: d, start: a.start, end: b.end}
}

func mergeBuf(a, b block) block {
	if a.start == b.end {
		a, b = b, a
	}
	buf := make([]int, b.end-a.start)

	ai, bi := 0, 0
	for i := 0; i < len(buf); i++ {
		if ai == len(a.data) || (bi < len(b.data) && a.data[ai] > b.data[bi]) {
			buf[i] = b.data[bi]
			bi++
		} else {
			buf[i] = a.data[ai]
			ai++
		}
	}

	d := a.data[0 : b.end-a.start]
	copy(d, buf)
	return block{data: d, start: a.start, end: b.end}
}

// findPair scans trough blocks slice and finds neighbour blocks
// returns found pair and true if a pair exists and empty pair and false otherwise
func findPair(blocks map[int]block) (pair, bool) {
	for _, v := range blocks {
		if b, ok := blocks[v.end]; ok {
			return pair{v, b}, true
		}
	}
	return pair{}, false
}

// rm removed block d from blocks slice (based on block's bounds)
func rm(blocks []block, d block) []block {
	for i, b := range blocks {
		if b.start == d.start && b.end == d.end {
			return append(blocks[:i], blocks[i+1:]...)
		}
	}
	return blocks
}

// a pair of blocks
type pair struct {
	a, b block
}

// block describes part of slice
// start and end are indices within initial slice
type block struct {
	start, end int
	data       []int
}

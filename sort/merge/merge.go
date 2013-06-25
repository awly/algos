// Mergesort
// for 1e6 elements
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
	mergeSort(data)
	fmt.Println(time.Since(start))
}

func mergeSort(data []int) {
	if len(data) < 2 {
		return
	}
	mergeSort(data[:len(data)/2])
	mergeSort(data[len(data)/2:])
	copy(data, merge(data[:len(data)/2], data[len(data)/2:]))
}

func merge(a []int, b []int) []int {
	res := make([]int, len(a)+len(b))
	ai, bi, i := 0, 0, 0
	for ; i < len(res); i++ {
		if a[ai] < b[bi] {
			res[i] = a[ai]
			ai++
			if ai == len(a) {
				break
			}
		} else {
			res[i] = b[bi]
			bi++
			if bi == len(b) {
				break
			}
		}
	}
	copy(res[i+1:], a[ai:])
	copy(res[i+1:], b[bi:])
	return res
}

// Quicksort
// on 1e7 slice
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	data := make([]int, 1e7)
	for i, _ := range data {
		data[i] = rand.Int()
	}
	start := time.Now()
	sort(data)
	fmt.Println(time.Since(start))
}

func sort(data []int) {
	if len(data) > 1 {
		pivot := rand.Intn(len(data)-1) + 1
		pivot = partition(data, pivot)
		sort(data[:pivot])
		sort(data[pivot+1:])
	}
}

func partition(data []int, pivot int) int {
	pivVal := data[pivot]
	data[len(data)-1], data[pivot] = data[pivot], data[len(data)-1]
	si := 0
	for i, v := range data[:len(data)-1] {
		if v < pivVal {
			data[i], data[si] = data[si], data[i]
			si++
		}
	}
	data[len(data)-1], data[si] = data[si], data[len(data)-1]
	return si
}

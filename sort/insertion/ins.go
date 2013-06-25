// Insertion sort
// for an array of 1e5 elements
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	data := make([]int, 1e5)
	for i := 0; i < len(data); i++ {
		data[i] = rand.Int()
	}
	start := time.Now()
	insSort(data)
	fmt.Println(time.Since(start))
}

func insSort(data []int) {
	pos := 0
	for key := 1; key < len(data); key++ {
		pos = binPosSearch(data[:key], data[key])
		for i := key; i > pos; i-- {
			data[i], data[i-1] = data[i-1], data[i]
		}
	}
}

func binPosSearch(data []int, val int) int {
	offset := 0
	for {
		if len(data) == 0 {
			break
		}
		if len(data) == 1 {
			if val >= data[0] {
				offset++
			}
			break
		}
		if val < data[len(data)/2] {
			data = data[:len(data)/2]
		} else {
			offset += len(data)/2 + 1
			data = data[len(data)/2+1:]
		}
	}
	return offset
}

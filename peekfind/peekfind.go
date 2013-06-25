// Peek find
// finds a "peek" in the given array
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
	findPeek(data)
	fmt.Println(time.Since(start))
}

func findPeek(data []int) int {
	switch len(data) {
	case 1:
		return data[0]
	case 2:
		if data[0] >= data[1] {
			return data[0]
		}
		return data[1]
	default:
		mid := int(len(data) / 2)
		if data[mid-1] >= data[mid] {
			return findPeek(data[:mid])
		} else if data[mid+1] >= data[mid] {
			return findPeek(data[mid+1:])
		}
		return data[mid]
	}
}

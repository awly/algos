// Parallel quicksort
// runs all sort calls in separate goroutines
// NumCPU active goroutines at max
// switches to insertion sort for slices smaller than the threshold
// sorts 1e7 random numbers in just over a second
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// chosen by measuring time for values 5..100
// not necessarily the best one...
var insSortThreshold = 36

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().Unix())

	wg := &sync.WaitGroup{}
	workers := make(chan struct{}, runtime.NumCPU())
	for i := 0; i < cap(workers); i++ {
		workers <- struct{}{}
	}
	data := make([]int, 1e7)
	for i, _ := range data {
		data[i] = rand.Int()
	}

	start := time.Now()
	wg.Add(1)
	sort(data, wg, workers)
	wg.Wait()
	fmt.Println(time.Since(start))
}

// parallel in-place quicksort
func sort(data []int, wg *sync.WaitGroup, workers chan struct{}) {
	<-workers
	if len(data) > 1 {
		pivot := rand.Intn(len(data)-1) + 1
		pivot = partition(data, pivot)
		if pivot > insSortThreshold {
			wg.Add(1)
			go sort(data[:pivot], wg, workers)
		} else {
			insSort(data[:pivot])
		}
		if len(data)-pivot > insSortThreshold {
			wg.Add(1)
			go sort(data[pivot+1:], wg, workers)
		} else {
			insSort(data[pivot+1:])
		}
	}
	workers <- struct{}{}
	wg.Done()
}

// move all elements > data[pivot] to the right and others to the left
// return new pivot index
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

// insertion sort for smaller slices
func insSort(data []int) {
	for key, pos := 1, 0; key < len(data); key++ {
		pos = binPosSearch(data[:key], data[key])
		for i := key; i > pos; i-- {
			data[i], data[i-1] = data[i-1], data[i]
		}
	}
}

// find position for the next insertion
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

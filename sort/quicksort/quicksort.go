// Quicksort
// normal and parallel versions
// both tested on 1e7 slice
// parallel version utilizes all cores
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

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
	parQuicksort(data, wg, workers, true)
	wg.Wait()
	fmt.Println("parallel:", time.Since(start))

	for i, _ := range data {
		data[i] = rand.Int()
	}
	start = time.Now()
	stdQuicksort(data)
	fmt.Println("normal:", time.Since(start))
}

func parQuicksort(data []int, wg *sync.WaitGroup, workers chan struct{}, isPar bool) {
	if isPar {
		<-workers
	}
	if len(data) > 1 {
		pivot := rand.Intn(len(data)-1) + 1
		pivot = partition(data, pivot)
		if pivot > 1000 {
			wg.Add(1)
			go parQuicksort(data[:pivot], wg, workers, true)
		} else {
			parQuicksort(data[:pivot], wg, workers, false)
		}
		if len(data)-pivot > 1000 {
			wg.Add(1)
			go parQuicksort(data[pivot+1:], wg, workers, true)
		} else {
			parQuicksort(data[pivot+1:], wg, workers, false)
		}
	}
	if isPar {
		workers <- struct{}{}
		wg.Done()
	}
}

func stdQuicksort(data []int) {
	if len(data) > 1 {
		pivot := rand.Intn(len(data)-1) + 1
		pivot = partition(data, pivot)
		stdQuicksort(data[:pivot])
		stdQuicksort(data[pivot+1:])
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

// Parallel quicksort
// runs all sort calls in separate goroutines
// GOMAXPROCS active goroutines at max
// switches to insertion sort for slices smaller than the threshold
// works best on large slices
package parquick

import (
	"github.com/captaincronos/algos/sort/insertion"
	"math/rand"
	"runtime"
	"sync"
)

// chosen by measuring time for values 5..100
// not necessarily the best one...
var insSortThreshold = 36

func Sort(data []int) {
	wg := &sync.WaitGroup{}
	workers := make(chan struct{}, runtime.GOMAXPROCS(0))
	for i := 0; i < cap(workers); i++ {
		workers <- struct{}{}
	}

	wg.Add(1)
	sort(data, wg, workers)
	wg.Wait()
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
			insertion.Sort(data[:pivot])
		}
		if len(data)-pivot > insSortThreshold {
			wg.Add(1)
			go sort(data[pivot+1:], wg, workers)
		} else {
			insertion.Sort(data[pivot+1:])
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

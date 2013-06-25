// Knapsack
// Solves knapsack problem for 1000 random items for a backpack of size 1000
// Uses dynamic programming, pseudo-polynomial time
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	items := make([]item, 1e3)
	for i, _ := range items {
		items[i] = item{rand.Intn(100), rand.Intn(100)}
	}
	selectItems(items, 1e3)
}

type item struct {
	size  int
	value int
}

func (i item) String() string {
	return fmt.Sprint("size:", i.size, " value:", i.value)
}

type state struct {
	value int
	incl  bool
}

func selectItems(items []item, size int) {
	memo := make([][]state, len(items))
	for i := len(items) - 1; i >= 0; i-- {
		memo[i] = make([]state, size+1)
		for left := 0; left <= size; left++ {
			if items[i].size > left {
				memo[i][left] = state{0, false}
				continue
			}
			if i == len(items)-1 {
				memo[i][left] = state{items[i].value, true}
				continue
			}
			if memo[i+1][left].value > (memo[i+1][left-items[i].size].value + items[i].value) {
				memo[i][left] = state{memo[i+1][left].value, false}
			} else {
				memo[i][left] = state{memo[i+1][left-items[i].size].value + items[i].value, true}
			}
		}
	}
	maxV, maxI := 0, 0
	for i := 0; i <= size; i++ {
		if memo[0][i].value > maxV {
			maxV = memo[0][i].value
			maxI = i
		}
	}
	// uncomment to see chosen items
	//left := maxI
	//for i := 0; i < len(items); i++ {
	//if memo[i][left].incl {
	//left -= items[i].size
	//fmt.Println(items[i])
	//}
	//}
	fmt.Println("max value:", maxV, "with size:", maxI)
}

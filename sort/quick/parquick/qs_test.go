package parquick

import (
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func randSlice(l int) []int {
	data := make([]int, l)
	for i := range data {
		data[i] = rand.Int()
	}
	return data
}

func BenchmarkParQuicksort100(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := randSlice(1e2)
		b.StartTimer()
		Sort(data)
		b.StopTimer()
	}
}

func BenchmarkParQuicksort1000(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := randSlice(1e3)
		b.StartTimer()
		Sort(data)
		b.StopTimer()
	}
}

func BenchmarkParQuicksort10000(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := randSlice(1e4)
		b.StartTimer()
		Sort(data)
		b.StopTimer()
	}
}

package pmerge

import (
	"math/rand"
	"testing"
)

func TestMergeEmpty(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Log("merge not panicked with wrong args")
			t.Fail()
		}
	}()
	merge(block{data: []int{}}, block{data: []int{}})
	// if it panics, it fails. nothing else to test
}

func TestMergeBadp1(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Log("merge not panicked with wrong args")
			t.Fail()
		}
	}()
	data := []int{1, 2, 3}
	merge(block{data: data}, block{start: -1, data: []int{}})
}

func TestMergeBadp2(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Log("merge not panicked with wrong args")
			t.Fail()
		}
	}()
	data := []int{1, 2, 3}
	merge(block{data: data}, block{start: len(data), data: []int{}})
}

func TestMergeNormal1(t *testing.T) {
	data := []int{1, 4, 6, 7, 2, 5, 6, 9}
	testMerge(t, data, 4)
	data = []int{1, 2, 3, 4, 6, 7, 8, 9}
	testMerge(t, data, 4)
	data = []int{6, 7, 8, 9, 1, 2, 3, 4}
	testMerge(t, data, 4)
}

func testMerge(t *testing.T, data []int, p int) {
	check := dup(data)
	merge(block{data: data[:p], end: p - 1}, block{start: p, end: len(data), data: data[p:]})
	if !sorted(data) {
		t.Logf("merge(%v, 4) -> %v\n", check, data)
		t.FailNow()
	}
}

func TestFindPairEmpty(t *testing.T) {
	_, ok := findPair(map[int]block{})
	if ok {
		t.Error("findPair returned ok for empty block slice")
	}
}

func TestFindPairNoPair(t *testing.T) {
	data := []int{1, 2, 3, 4, 6, 7, 8, 9}
	_, ok := findPair(map[int]block{
		0: block{start: 0, end: 1, data: data[0:2]},
		3: block{start: 3, end: 4, data: data[3:4]},
		6: block{start: 6, end: 7, data: data[6:7]},
	})
	if ok {
		t.Error("findPair returned ok for blocks with no overlaps")
	}
}

func TestFindPairNormal(t *testing.T) {
	data := []int{1, 2, 3, 4, 6, 7, 8, 9}
	testFindPair(t, map[int]block{
		0: block{start: 0, end: 2, data: data[0:2]},
		2: block{start: 2, end: 4, data: data[2:3]},
		4: block{start: 4, end: 6, data: data[4:5]},
		6: block{start: 6, end: 8, data: data[6:7]},
	})
	testFindPair(t, map[int]block{
		0: block{start: 0, end: 2, data: data[0:2]},
		4: block{start: 4, end: 6, data: data[4:5]},
		2: block{start: 2, end: 4, data: data[2:3]},
		6: block{start: 6, end: 8, data: data[6:7]},
	})
	testFindPair(t, map[int]block{
		0: block{start: 0, end: 2, data: data[0:2]},
		3: block{start: 3, end: 5, data: data[3:4]},
		5: block{start: 5, end: 8, data: data[5:7]},
	})
}

func testFindPair(t *testing.T, d map[int]block) {
	p, ok := findPair(d)
	if !ok {
		t.Log("findPair failed to find a pair")
		t.FailNow()
	}
	if p.a.start != p.b.end && p.b.start != p.a.end {
		t.Log("findPair returned not neighbouring blocks")
		t.FailNow()
	}
}

func TestSortEmpty(t *testing.T) {
	data := []int{}
	Sort(data)
	// should not panic
}

func TestSortNormal(t *testing.T) {
	data := rand.Perm(100)
	Sort(data)
	if !sorted(data) {
		t.Error("Sort returned unsorted data")
	}
}

func randSlice(l int) []int {
	data := make([]int, l)
	for i := range data {
		data[i] = rand.Int()
	}
	return data
}

func BenchmarkSort10(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := randSlice(1e1)
		b.StartTimer()
		Sort(data)
		b.StopTimer()
	}
}

func BenchmarkSort100(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := randSlice(1e2)
		b.StartTimer()
		Sort(data)
		b.StopTimer()
	}
}

func BenchmarkSort1000(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := randSlice(1e3)
		b.StartTimer()
		Sort(data)
		b.StopTimer()
	}
}

func sorted(a []int) bool {
	for i := 0; i < len(a)-2; i++ {
		if a[i] > a[i+1] {
			return false
		}
	}
	return true
}

func dup(a []int) []int {
	b := make([]int, len(a))
	copy(b, a)
	return b
}

func eq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

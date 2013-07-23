package orset

import (
	"testing"
)

func TestSlow(t *testing.T) {
	test(NewSlow(), t)
}

func TestFast(t *testing.T) {
	test(NewFast(), t)
}

func test(set Orset, t *testing.T) {
	set.PushBack(Int(1))
	if set.PopFront() != Int(1) {
		t.Fatal("popped element does not match pushed one")
	}
	for i := 0; i < 10; i++ {
		set.PushBack(Int(i))
	}
	for i := 0; i < 10; i++ {
		if set.PopFront() != Int(i) {
			t.Fatal("popped element does not match pushed one")
		}
	}
	if set.PopFront() != nil {
		t.Fatal("popped non-nil element from empty queue")
	}
}

func BenchmarkSlowPush(b *testing.B) {
	benchPush(NewSlow(), b)
}

func BenchmarkFastPush(b *testing.B) {
	benchPush(NewFast(), b)
}

func benchPush(set Orset, b *testing.B) {
	for i := 0; i < b.N; i++ {
		set.PushBack(Int(i))
	}
}

func BenchmarkSlowPushDup(b *testing.B) {
	benchPushDup(NewSlow(), b)
}

func BenchmarkFastPushDup(b *testing.B) {
	benchPushDup(NewFast(), b)
}

func benchPushDup(set Orset, b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		set.PushBack(Int(i))
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		set.PushBack(Int(i))
	}
}

// for some reason this runs too long and gets killed by go test. don't run this.
func BenchmarkSlowPop(b *testing.B) {
	benchPop(NewSlow(), b)
}

func BenchmarkFastPop(b *testing.B) {
	benchPop(NewFast(), b)
}

func benchPop(set Orset, b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		set.PushBack(Int(i))
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		set.PopFront()
	}
}

package ringbuf

import (
	"fmt"
	"slices"
	"testing"
)

func ensure(cond bool, msg string, args ...any) {
	if !cond {
		panic(fmt.Sprintf(msg, args...))
	}
}

func TestRingbufEmptyLength(t *testing.T) {
	r := NewRingBuf[int](5)

	want := 0
	have := r.Len()

	ensure(have == want, "have: %d want: %d", have, want)
}

func TestRingbufEmptyCapacity(t *testing.T) {
	r := NewRingBuf[int](5)

	want := 5
	have := r.Cap()

	ensure(have == want, "have: %d want: %d", have, want)
}

func TestRingbufEmptyIterNext(t *testing.T) {
	r := NewRingBuf[int](5)
	i := r.Iter()

	want := false
	have := i.Next()

	ensure(have == want, "have: %t want: %t", have, want)
}

func TestRingbufLength(t *testing.T) {
	r := NewRingBuf[int](5)
	r.Put(11, 22)

	want := 2
	have := r.Len()

	ensure(have == want, "have: %d want: %d", have, want)
}

func TestRingbufCapacity(t *testing.T) {
	r := NewRingBuf[int](5)
	r.Put(11, 22, 33)

	want := 5
	have := r.Cap()

	ensure(have == want, "have: %d want: %d", have, want)
}

func TestRingbufIterNext(t *testing.T) {
	r := NewRingBuf[int](5)
	r.Put(11, 22)
	i := r.Iter()

	want := true
	have := i.Next()

	ensure(have == want, "have: %t want: %t", have, want)
}

func TestRingbufIterGet(t *testing.T) {
	r := NewRingBuf[int](5)
	r.Put(11, 22, 33)
	i := r.Iter()

	want := []int{11, 22, 33}
	have := []int{}
	for i.Next() {
		have = append(have, *i.Get())
	}

	ensure(slices.Compare(have, want) == 0, "have: %v want: %v", have, want)
}

func TestRingbufRaw(t *testing.T) {
	r := NewRingBuf[int](5)
	r.Put(11, 22, 33)

	want := []int{11, 22, 33}
	have := r.Raw()

	ensure(slices.Compare(have, want) == 0, "have: %v want: %v", have, want)
}

func TestRingbufFullLength(t *testing.T) {
	r := NewRingBuf[int](3)
	r.Put(11, 22, 33, 44)

	want := 3
	have := r.Len()

	ensure(have == want, "have: %d want: %d", have, want)
}

func TestRingbufFullCapacity(t *testing.T) {
	r := NewRingBuf[int](3)
	r.Put(11, 22, 33, 44)

	want := 3
	have := r.Cap()

	ensure(have == want, "have: %d want: %d", have, want)
}

func TestRingbufFullIter(t *testing.T) {
	r := NewRingBuf[int](5)
	r.Put(11, 22, 33, 44, 55, 66, 77)
	i := r.Iter()

	want := []int{33, 44, 55, 66, 77}
	have := []int{}
	for i.Next() {
		have = append(have, *i.Get())
	}

	ensure(slices.Compare(have, want) == 0, "have: %v want: %v", have, want)
}

func TestRingbufFullRaw(t *testing.T) {
	r := NewRingBuf[int](5)
	r.Put(11, 22, 33, 44, 55, 66, 77)

	want := []int{66, 77, 33, 44, 55}
	have := r.Raw()

	ensure(slices.Compare(have, want) == 0, "have: %v want: %v", have, want)
}

package problem1

import (
	"container/heap"
	"testing"
)

func TestCorrectness(t *testing.T) {
	h := &IntHeap{}
	heap.Init(h)
	heap.Push(h, 4)
	heap.Push(h, 1)
	heap.Push(h, 100)
	heap.Push(h, 99)
	heap.Push(h, 5)
	heap.Push(h, 3)

	expected := []int{100, 99, 5, 4, 3, 1}
	for i := 0; i < h.Len(); i++ {
		current := heap.Pop(h)
		if current != expected[i] {
			t.Errorf("%d - Got %d expected %d", i, current, expected[i])
		}
	}

}

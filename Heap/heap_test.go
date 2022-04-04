package heap

import (
	"testing"
)

func TestMaxHeap_Insert(t *testing.T) {
	m := &MaxHeap{}
	buildHeap := []int{10, 20, 30, 40, 50, 60, 60}

	for _, v := range buildHeap {
		m.Insert(v)
	}

	insertExpected := []int{60, 40, 60, 10, 30, 20, 50}
	for i, v := range m.array {
		if v != insertExpected[i] {
			t.Fatalf("value: %d, expected: %d failed", v, insertExpected[i])
		}

		t.Logf("value: %d, expected: %d success", v, insertExpected[i])
	}
}

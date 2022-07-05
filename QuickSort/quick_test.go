package QuickSort

import "testing"

func TestQuickSort(t *testing.T) {
	s := []int{3, 5, 10, 1, 7, 2}
	QuickSort(s, 0, len(s)-1)
	t.Log(s)
}

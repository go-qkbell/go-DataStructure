package MergeSort

import "testing"

func TestMergeSort(t *testing.T) {
	s := []int{3, 5, 10, 1, 7, 2}
	res := MergeSort(s)
	t.Log(res)
}

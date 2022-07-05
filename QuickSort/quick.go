package QuickSort

func QuickSort(nums []int, start, end int) {
	if start < end {
		pivotIndex := Partition(nums, start, end)

		QuickSort(nums, start, pivotIndex-1)
		QuickSort(nums, pivotIndex+1, end)
	}
}

func Partition(nums []int, start, end int) int {
	pivot := nums[end]
	pivotIndex := start

	for i := start; i < end; i++ {
		if nums[i] < pivot {
			nums[pivotIndex], nums[i] = nums[i], nums[pivotIndex]
			pivotIndex++
		}
	}

	nums[pivotIndex], nums[end] = nums[end], nums[pivotIndex]

	return pivotIndex
}

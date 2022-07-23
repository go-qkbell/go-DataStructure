package MergeSort

func Merge(left, right []int) []int {
	var res []int
	var i, j int

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
		}
	}

	for ; i < len(left); i++ {
		res = append(res, left[i])
	}

	for ; j < len(right); j++ {
		res = append(res, right[j])
	}

	return res
}

func MergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	mid := len(nums) / 2

	left := MergeSort(nums[:mid])
	right := MergeSort(nums[mid:])

	res := Merge(left, right)

	return res
}

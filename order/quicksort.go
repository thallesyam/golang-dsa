package order

func Quicksort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	pivot := nums[0]
	var less []int
	var bigger []int

	for i := 1; i < len(nums); i++ {
		if nums[i] < pivot {
			less = append(less, nums[i])
		}

		if nums[i] > pivot {
			bigger = append(bigger, nums[i])
		}
	}

	return append(append(Quicksort(less), pivot), Quicksort(bigger)...)
}

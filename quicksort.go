package quicksort

// QuickSort 对传入的整数切片进行原地快速排序。
// 算法平均时间复杂度为 O(n log n)，最坏情况为 O(n^2)。
// 通过递归地划分切片并分别排序子切片实现。
func QuickSort(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}

	pivotIndex := partition(nums, left, right)
	quickSort(nums, left, pivotIndex-1)
	quickSort(nums, pivotIndex+1, right)
}

func partition(nums []int, left, right int) int {
	pivot := nums[right]
	i := left

	for j := left; j < right; j++ {
		if nums[j] <= pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}

	nums[i], nums[right] = nums[right], nums[i]
	return i
}

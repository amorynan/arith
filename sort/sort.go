package sort

func DistributionAndMerge(nums []int) []int {
	return DistAndMergeLoop(nums)
}

func DistAndMergeLoop(nums []int) []int {
	mid := len(nums) >> 1
	right := nums[:mid]
	left := nums[mid:]

	if len(right) == 1 && len(left) == 1 {
		if right[0] >= left[0] {
			return []int{left[0], right[0]}
		}
		return []int{right[0], left[0]}
	}

	sortedRight,  sortedLeft := []int{}, []int{}
	if len(right) == 1 {
		// here need left cal
		sortedLeft = DistAndMergeLoop(left)
		sortedRight = []int{right[0]}
	}else if len(left) == 1 {
		sortedLeft = []int{left[0]}
		sortedRight = DistAndMergeLoop(right)
	}else {
		sortedRight = DistAndMergeLoop(right)
		sortedLeft  = DistAndMergeLoop(left)
	}

	sorted := make([]int, 0 ,len(sortedLeft)+len(sortedRight))
	ri, ji := 0, 0
	for ri < len(sortedRight) && ji < len(sortedLeft) {
		if sortedRight[ri] >= sortedLeft[ji] {
			sorted = append(sorted, sortedLeft[ji])
			ji++
		}else {
			sorted = append(sorted, sortedRight[ri])
			ri ++
		}
	}
	if ri < len(sortedRight) {
		sorted = append(sorted, sortedRight[ri:]...)
	}
	if ji < len(sortedLeft) {
		sorted = append(sorted, sortedLeft[ji:]...)
	}
	return sorted
}
package array

import "math"

func twoSum(nums []int, target int) []int {
	// only go through once, because first just can sum with the left in last
	leftMap := make(map[int]int, len(nums))
	for idx, item := range nums {
		if val, exist := leftMap[target - nums[idx]]; exist {
			return []int{idx, val}
		}
		leftMap[item] = idx
	}
	return nil
}


func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	mid := (len(nums1) + len(nums2)) / 2
	midLeft := -1
	// odd ?
	if (len(nums1) + len(nums2)) % 2 == 0 {
		midLeft = mid-1
	}
	totalNum := make([]int, 0, len(nums1)+len(nums2))
	idx1, idx2 := 0, 0
	for idx:=0; idx< cap(totalNum); idx++ {
		if idx1 >= len(nums1) {
			totalNum = append(totalNum, nums2[idx2:]...)
			break
		}

		if idx2 >= len(nums2) {
			totalNum = append(totalNum, nums1[idx1:]...)
			break
		}

		if nums1[idx1] < nums2[idx2]{
			totalNum = append(totalNum, nums1[idx1])
			idx1 ++
		}else {
			totalNum = append(totalNum, nums2[idx2])
			idx2 ++
		}
	}

	if midLeft < 0 {
		return float64(totalNum[mid])
	}else {
		return float64(totalNum[midLeft] + totalNum[mid]) / 2.0
	}
}